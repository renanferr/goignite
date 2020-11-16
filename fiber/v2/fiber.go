package gifiber

import (
	"context"
	"strconv"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/gofiber/fiber/v2"
)

var (
	app *fiber.App
)

const (
	TopicApp = "topic:giecho:app"
)

func Start(ctx context.Context) *fiber.App {
	options, _ := AppConfig()
	app = fiber.New(*options)

	gieventbus.Publish(TopicApp, app)

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
