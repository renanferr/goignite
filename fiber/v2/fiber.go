package gifiber

import (
	"context"
	"strconv"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
)

type Ext func(context.Context, *fiber.App) error

func Start(ctx context.Context, exts ...Ext) *fiber.App {

	config, _ := AppConfig()
	app = fiber.New(*config)

	for _, ext := range exts {
		if err := ext(ctx, app); err != nil {
			panic(err)
		}
	}

	return app
}

func Serve(ctx context.Context) {

	logger := gilog.FromContext(ctx)
	logger.Infof("starting fiber server. https://gofiber.io/")

	logger.Fatal(app.Listen(serverPort()))
}

func serverPort() string {
	return ":" + strconv.Itoa(Port())
}
