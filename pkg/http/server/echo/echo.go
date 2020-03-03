package echo

import (
	"context"
	"strconv"

	srvconfig "github.com/b2wdigital/goignite/pkg/http/server/config"
	"github.com/b2wdigital/goignite/pkg/http/server/echo/config"
	"github.com/b2wdigital/goignite/pkg/http/server/echo/handler"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	m "github.com/neko-neko/echo-logrus/v2"
	"github.com/neko-neko/echo-logrus/v2/log"
)

var (
	instance *echo.Echo
)

func Start(ctx context.Context) *echo.Echo {

	instance = echo.New()

	instance.HideBanner = config.GetHideBanner()
	instance.Logger = log.Logger()

	setDefaultMiddlewares(ctx, instance)
	setDefaultRouters(ctx, instance)

	return instance
}

func setDefaultMiddlewares(ctx context.Context, instance *echo.Echo) {
	instance.Use(m.Logger())
	instance.Use(middleware.Recover())
}

func setDefaultRouters(ctx context.Context, instance *echo.Echo) {

	l := logrus.FromContext(ctx)

	statusRoute := srvconfig.GetStatusRoute()

	l.Infof("configuring status router on %s", statusRoute)

	statusHandler := handler.NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

	healthRoute := srvconfig.GetHealthRoute()

	l.Infof("configuring health router on %s", healthRoute)

	healthHandler := handler.NewHealthHandler()
	instance.GET(healthRoute, healthHandler.Get)
}

func Serve(ctx context.Context) {
	l := logrus.FromContext(ctx)
	l.Info("starting echo server. https://echo.labstack.com/")
	instance.Logger.Fatal(instance.Start(getServerPort()))
}

func getServerPort() string {
	return ":" + strconv.Itoa(srvconfig.GetPort())
}
