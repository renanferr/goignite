package config

import (
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/jpfaria/goignite/pkg/config"
)

const (
	Port        = "serverless.cloudnative.http.port"
	Path        = "serverless.cloudnative.http.path"
	ContentType = "serverless.cloudnative.http.contenttype"
)

func init() {

	log.Println("getting configurations for http server")

	config.Add(Port, 8080, "define http port")
	config.Add(Path, "/", "define path")
	config.Add(ContentType, cloudevents.ApplicationJSON, "define content type")
}

func GetPort() int {
	return config.Int(Port)
}

func GetPath() string {
	return config.String(Path)
}

func GetContentType() string {
	return config.String(ContentType)
}
