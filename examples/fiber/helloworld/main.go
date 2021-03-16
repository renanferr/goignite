package main

import (
	"context"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
	gifibercors "github.com/b2wdigital/goignite/v2/fiber/v2/ext/cors"
	gifiberetag "github.com/b2wdigital/goignite/v2/fiber/v2/ext/etag"
	giinfo "github.com/b2wdigital/goignite/v2/info"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
	giserver "github.com/b2wdigital/goignite/v2/server"
	"github.com/gofiber/fiber/v2"
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

func Get(c *fiber.Ctx) (err error) {

	logger := gilog.FromContext(context.Background())

	resp := Response{
		Message: "Hello World!!",
	}

	err = giconfig.Unmarshal(&resp)
	if err != nil {
		logger.Errorf(err.Error())
	}

	return c.Status(http.StatusOK).JSON(resp)
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

	giinfo.AppName = "helloworld"

	fiberSrv := gifiber.NewDefault(ctx,
		gifibercors.Register,
		gifiberetag.Register)

	fiberSrv.App().Get(c.App.Endpoint.Helloworld, Get)

	fiberSrv.Serve(ctx)

	chiSrv := gichi.NewDefault(ctx)

	giserver.Serve(ctx, fiberSrv, chiSrv)
}
