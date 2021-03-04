package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/cors"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/gzip"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/health"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/logger"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/requestid"
	"github.com/b2wdigital/goignite/v2/echo/v4/ext/status"
	"github.com/b2wdigital/goignite/v2/info"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
	"github.com/wesovilabs/beyond/api"
)

const Endpoint = "app.endpoint.google"

func init() {
	giconfig.Add(Endpoint, "/google", "google endpoint")
}

func Beyond() *api.Beyond {
	return api.New().
		WithBefore(NewTracingAdvice, "handler.Get(...)").
		WithBefore(NewTracingAdviceWithPrefix("[beyond]"), "handler.*(...)...")
}

func main() {

	var err error

	giconfig.Load()

	c := Config{}

	err = giconfig.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	gilogrus.NewLogger()

	info.AppName = "google"

	instance := giecho.New(ctx,
		cors.Register,
		requestid.Register,
		gzip.Register,
		logger.Register,
		status.Register,
		health.Register)

	instance.GET(c.App.Endpoint.Google, Get)

	giecho.Serve(ctx)
}
