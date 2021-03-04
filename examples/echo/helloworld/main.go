package main

import (
	"context"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/cors"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/gzip"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/health"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/logger"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/requestid"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/status"
	"github.com/b2wdigital/goignite/v2/info"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
	"github.com/labstack/echo/v4"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"

func init() {
	giconfig.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
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

func Get(c echo.Context) (err error) {

	logger := gilog.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = giconfig.Unmarshal(&resp)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return giecho.JSON(c, http.StatusOK, resp, err)
}

func main() {

	giconfig.Load()

	c := Config{}

	err := giconfig.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	gilogrus.NewLogger()

	info.AppName = "helloworld"

	instance := giecho.New(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		logger.Register,
		status.Register,
		health.Register)

	instance.GET(c.App.Endpoint.Helloworld, Get)

	giecho.Serve(ctx)
}
