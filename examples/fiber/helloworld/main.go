package main

import (
	"context"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
	"github.com/b2wdigital/goignite/v2/fiber/v2/ext/cors"
	"github.com/b2wdigital/goignite/v2/fiber/v2/ext/etag"
	"github.com/b2wdigital/goignite/v2/info"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
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

	info.AppName = "helloworld"

	instance := gifiber.Start(ctx,
		cors.Register,
		etag.Register)

	instance.Get(c.App.Endpoint.Helloworld, Get)

	gifiber.Serve(ctx)
}
