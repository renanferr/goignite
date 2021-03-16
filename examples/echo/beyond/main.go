package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
	giechocors "github.com/b2wdigital/goignite/v2/echo/v4/ext/cors"
	giechogzip "github.com/b2wdigital/goignite/v2/echo/v4/ext/gzip"
	giechohealth "github.com/b2wdigital/goignite/v2/echo/v4/ext/health"
	giechologger "github.com/b2wdigital/goignite/v2/echo/v4/ext/logger"
	giechorequestid "github.com/b2wdigital/goignite/v2/echo/v4/ext/requestid"
	giechostatus "github.com/b2wdigital/goignite/v2/echo/v4/ext/status"
	giinfo "github.com/b2wdigital/goignite/v2/info"
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

	giinfo.AppName = "google"

	srv := giecho.NewDefault(ctx,
		giechocors.Register,
		giechorequestid.Register,
		giechogzip.Register,
		giechologger.Register,
		giechostatus.Register,
		giechohealth.Register)

	srv.Echo().GET(c.App.Endpoint.Google, Get)

	srv.Serve(ctx)
}
