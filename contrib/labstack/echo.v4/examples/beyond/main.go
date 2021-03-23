package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/core/health"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/core/log"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/core/status"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/cors"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/gzip"
	"github.com/b2wdigital/goignite/v2/contrib/labstack/echo.v4/ext/requestid"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	"github.com/b2wdigital/goignite/v2/core/info"
	"github.com/wesovilabs/beyond/api"
)

const Endpoint = "app.endpoint.google"

func init() {
	config.Add(Endpoint, "/google", "google endpoint")
}

func Beyond() *api.Beyond {
	return api.New().
		WithBefore(NewTracingAdvice, "handler.Get(...)").
		WithBefore(NewTracingAdviceWithPrefix("[beyond]"), "handler.*(...)...")
}

func main() {

	var err error

	config.Load()

	c := Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	logrus.NewLogger()

	info.AppName = "google"

	srv := echo.NewDefault(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		log.Register,
		status.Register,
		health.Register)

	srv.Echo().GET(c.App.Endpoint.Google, Get)

	srv.Serve(ctx)
}
