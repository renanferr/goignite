package echo

import (
	"context"
	"strconv"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/server/http/rest"
	"github.com/labstack/echo/v4/middleware"
	m "github.com/neko-neko/echo-logrus/v2"
	elog "github.com/neko-neko/echo-logrus/v2/log"
)

var (
	instance *echo.Echo
)

func Start(ctx context.Context) *echo.Echo {

	instance = echo.New()

	instance.HideBanner = GetHideBanner()
	instance.Logger = elog.Logger()

	setDefaultMiddlewares(ctx, instance)
	setDefaultRouters(ctx, instance)

	return instance
}

func setDefaultMiddlewares(ctx context.Context, instance *echo.Echo) {
	instance.Use(m.Logger())
	instance.Use(middleware.Recover())
}

func setDefaultRouters(ctx context.Context, instance *echo.Echo) {

	l := log.FromContext(ctx)

	statusRoute := rest.GetStatusRoute()

	l.Infof("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

	healthRoute := rest.GetHealthRoute()

	l.Infof("configuring health router on %s", healthRoute)

	healthHandler := NewHealthHandler()
	instance.GET(healthRoute, healthHandler.Get)
}

func Serve(ctx context.Context) {
	l := log.FromContext(ctx)
	l.Infof("starting echo server. https://echo.labstack.com/")
	err := instance.Start(getServerPort())
	instance.Logger.Fatalf(err.Error())
}

func getServerPort() string {
	return ":" + strconv.Itoa(rest.GetPort())
}
