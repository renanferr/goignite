package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/echo/v4"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/cors"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/gzip"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/health"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/logger"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/requestid"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/status"
	"github.com/b2wdigital/goignite/v2/info"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
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
		logger.Register,
		status.Register,
		health.Register)

	srv.Echo().GET(c.App.Endpoint.Google, Get)

	srv.Serve(ctx)
}
