package gifiber

import (
	"context"
	"strconv"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
)

func Start(ctx context.Context, exts ...func(context.Context, *fiber.App) error) *fiber.App {

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

	l := gilog.FromContext(ctx)
	l.Infof("starting fiber server. https://gofiber.io/")

	l.Fatal(app.Listen(serverPort()))
}

func serverPort() string {
	return ":" + strconv.Itoa(Port())
}
