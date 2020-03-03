package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const (
	Port = "serverless.cloudnative.http.port"
	Path = "serverless.cloudnative.http.path"
)

func init() {

	log.Println("getting configurations for http server")

	config.Add(Port, 8080, "define http port")
	config.Add(Path, "/", "define path")
}

func GetPort() int {
	return config.Int(Port)
}

func GetPath() string {
	return config.String(Path)
}
