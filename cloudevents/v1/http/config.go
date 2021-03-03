package http

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	cloudevents "github.com/cloudevents/sdk-go"
)

const (
	Port        = "gi.cloudevents.http.port"
	Path        = "gi.cloudevents.http.path"
	ContentType = "gi.cloudevents.http.contentType"
)

func init() {

	giconfig.Add(Port, 8080, "define http port")
	giconfig.Add(Path, "/", "define path")
	giconfig.Add(ContentType, cloudevents.ApplicationJSON, "define content type")
}

func GetPort() int {
	return giconfig.Int(Port)
}

func GetPath() string {
	return giconfig.String(Path)
}

func GetContentType() string {
	return giconfig.String(ContentType)
}
