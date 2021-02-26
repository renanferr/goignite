package gifiber

import (
	"context"
	"strconv"
	"strings"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var (
	app *fiber.App
)

const (
	TopicApp = "topic:gifiber:app"
)

func Start(ctx context.Context) *fiber.App {

	config, _ := AppConfig()
	app = fiber.New(*config)

	setDefaultRouters(ctx, app)
	setDefaultMiddlewares(app)

	gieventbus.Publish(TopicApp, app)

	return app
}

func setDefaultMiddlewares(app *fiber.App) {

	if GetMiddlewareRecoverEnabled() {
		app.Use(recover.New())
	}

	if GetMiddlewareLogEnabled() {
		app.Use(logger.New())
	}

	if GetMiddlewareCORSEnabled() {

		app.Use(cors.New(cors.Config{
			AllowOrigins:     strings.Join(GetMiddlewareCORSAllowOrigins(), ","),
			AllowMethods:     strings.Join(GetMiddlewareCORSAllowMethods(), ","),
			AllowHeaders:     strings.Join(GetMiddlewareCORSAllowHeaders(), ","),
			AllowCredentials: GetMiddlewareCORSAllowCredentials(),
			ExposeHeaders:    strings.Join(GetMiddlewareCORSExposeHeaders(), ","),
			MaxAge:           GetMiddlewareCORSMaxAge(),
		}))
	}

	if GetMiddlewareRequestIDEnabled() {
		app.Use(requestid.New())
	}

	if GetMiddlewarePprofEnabled() {
		app.Use(pprof.New())
	}

	if GetMiddlewareETagEnabled() {
		app.Use(etag.New())
	}

	if GetMiddlewareMonitorEnabled() {
		app.Use(monitor.New())
	}

}

func setDefaultRouters(ctx context.Context, app *fiber.App) {

	l := gilog.FromContext(ctx)

	statusRoute := GetStatusRoute()

	l.Infof("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()
	app.Get(statusRoute, statusHandler.Get)

	healthRoute := GetHealthRoute()

	l.Infof("configuring health router on %s", healthRoute)

	healthHandler := NewHealthHandler()
	app.Get(healthRoute, healthHandler.Get)
}

func Serve(ctx context.Context) {

	l := gilog.FromContext(ctx)
	l.Infof("starting fiber server. https://gofiber.io/")

	l.Fatal(app.Listen(serverPort()))
}

func serverPort() string {
	return ":" + strconv.Itoa(Port())
}
