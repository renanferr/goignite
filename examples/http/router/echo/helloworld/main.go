package main

import (
	"context"
	"log"
	"net/http"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/http/router/echo"
	"github.com/b2wdigital/goignite/pkg/http/router/echo/parser"
	"github.com/b2wdigital/goignite/pkg/info"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
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

	l := logrus.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Error(err)
	}

	return parser.JSONResponse(c, http.StatusOK, resp, err)
}

func main() {

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	c := Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	logrus.Start()

	info.AppName = "helloworld"

	instance := echo.Start(ctx)

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	instance.GET(c.App.Endpoint.Helloworld, Get)

	echo.Serve(ctx)
}
