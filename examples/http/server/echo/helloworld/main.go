package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/http/server/echo"
	"github.com/jpfaria/goignite/pkg/http/server/echo/parser"
	"github.com/jpfaria/goignite/pkg/info"
	"github.com/jpfaria/goignite/pkg/log/logrus"
	e "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	c := Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}

	logrus.Start()

	info.AppName = "helloworld"

	instance := echo.Start()

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	instance.GET(c.App.Endpoint.Helloworld, Get)

	echo.Serve()
}
