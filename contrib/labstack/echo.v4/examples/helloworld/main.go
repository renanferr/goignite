package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/core/health"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/core/log"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/core/status"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/cors"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/gzip"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/requestid"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/info"
	"github.com/b2wdigital/goignite/v2/core/log"
	e "github.com/labstack/echo/v4"
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

	logger := log.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return echo.JSON(c, http.StatusOK, resp, err)
}

func main() {

	config.Load()

	c := Config{}

	err := config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	//logrus.NewLogger()
	//zap.NewLogger()
	//zap.log.NewLogger()

	info.AppName = "helloworld"

	srv := echo.NewDefault(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		log.Register,
		status.Register,
		health.Register)

	srv.Echo().GET(c.App.Endpoint.Helloworld, Get)

	srv.Serve(ctx)
}
