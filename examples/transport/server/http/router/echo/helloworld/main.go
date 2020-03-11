package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/info"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/server/http/router/echo/v4"
	e "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"

func init() {
	config.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
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

	l := log.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Errorf(err.Error())
	}

	return v4.JSONResponse(c, http.StatusOK, resp, err)
}

func main() {

	err := config.Load()
	if err != nil {
		panic(err)
	}

	c := Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	log.NewLogger(v1.NewLogger())

	info.AppName = "helloworld"

	instance := v4.Start(ctx)

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	instance.GET(c.App.Endpoint.Helloworld, Get)

	v4.Serve(ctx)
}
