package main

import (
	"context"
    "log"
    "net/http"

    "github.com/jpfaria/goignite/pkg/config"
    "github.com/jpfaria/goignite/pkg/http/server/echo"
    "github.com/jpfaria/goignite/pkg/http/server/echo/parser"
    "github.com/jpfaria/goignite/pkg/info"
    "github.com/jpfaria/goignite/pkg/logging/logrus"
    e "github.com/labstack/echo"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"
const ResponseMessage = "message"

func init() {
	config.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
	config.Add(ResponseMessage, "hello world!!!", "default response message")
}

type Config struct {
	App struct {
		Endpoint struct {
		    Helloworld string
		}
	}
}

type Response struct {
	Message string
}

func Get(c e.Context) (err error) {

	l := logrus.FromContext(context.Background())

	resp := Response{}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Error(err)
	}

	return parser.JSONResponse(c, http.StatusOK, resp, err)
}

func main() {

	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	c := Config{}

    err = config.Unmarshal(&c)
    if err != nil {
        log.Fatal(err)
    }

	info.AppName = "helloworld"

	instance := echo.Start()

	instance.GET(c.App.Endpoint.Helloworld, Get)

	echo.Serve()
}
