package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"text/template"

	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/log/logrus"
)

type Data struct {
	RequestMaps []*RequestMap
	Packages    []*Package
}

type Package struct {
	Alias string
	URI   string
}

type RequestMap struct {
	Method   string
	Endpoint string
	Handler  *Handler
	Body     *Body
}

type Handler struct {
	Package string
	Func    string
	Alias   string
}

type Body struct {
	Package string
	Struct  string
	Alias   string
}

func main() {

	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	logrus.Start()

	tmpl := template.Must(template.ParseFiles("templates/http/server/echo.tpl"))

	data := Data{
		RequestMaps: []*RequestMap{
			{
				Method:   "GET",
				Endpoint: "/test1/:test",
				Handler: &Handler{
					Package: "github.com/jpfaria/goignite/pkg/serverless/http/server/example/handler",
					Func:    "Test1",
				},
			},
			{
				Method:   "POST",
				Endpoint: "/test2",
				Handler: &Handler{
					Package: "github.com/jpfaria/goignite/pkg/serverless/http/server/example/handler",
					Func:    "Test2",
				},
				Body: &Body{
					Package: "github.com/jpfaria/goignite/pkg/serverless/http/server/example/model/request",
					Struct: "User",
				},
			},
		},
	}

	for _, r := range data.RequestMaps {
		r.Handler.Alias = getAlias(r.Handler.Package)

		if r.Method == "POST" || r.Method == "PUT" || r.Method == "PATCH" {
			r.Body.Alias = getAlias(r.Body.Package)
		}
	}

	data.Packages = getPackages(data.RequestMaps)

	var buf bytes.Buffer

	err = tmpl.Execute(&buf, data)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("gen/main.go", buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func getPackages(maps []*RequestMap) []*Package {
	var packages []string

	for _, v := range maps {
		packages = append(packages, v.Handler.Package)
		if v.Body != nil {
			packages = append(packages, v.Body.Package)
		}
	}

	packages = unique(packages)

	var p []*Package

	for _, uri := range packages {
		alias := getAlias(uri)
		p = append(p, &Package{Alias: alias, URI: uri})
	}

	return p
}

func getAlias(uri string) string {
	hasher := md5.New()
	hasher.Write([]byte(uri))
	return hex.EncodeToString(hasher.Sum(nil))
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
