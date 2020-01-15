package main

import (
    "github.com/jpfaria/goignite/pkg/config"
    "github.com/jpfaria/goignite/pkg/http/server/echo"
    "github.com/jpfaria/goignite/pkg/info"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"
const ResponseMessage = "message"

func init() {
    config.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
    config.Add(ResponseMessage, "hello world!!!", "default response message")
}

func main() {

    config.Parse()

    info.AppName = "rest_server_echo"

    instance := echo.Start()

    handler := NewHandler()
    instance.GET(config.Instance.String(HelloWorldEndpoint), handler.Get)

    echo.Serve()
}
