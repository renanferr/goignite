package echo

import (
	"context"
	"strconv"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/client/newrelic/v3"
	mware "github.com/b2wdigital/goignite/pkg/transport/server/echo/v4/middleware"
	prometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	instance *echo.Echo
)

func Start(ctx context.Context) *echo.Echo {

	instance = echo.New()

	instance.HideBanner = GetHideBanner()
	instance.Logger = Wrap(log.GetLogger())

	setDefaultMiddlewares(instance)
	setDefaultRouters(ctx, instance)

	return instance
}

func setDefaultMiddlewares(instance *echo.Echo) {

	if GetMiddlewareLogEnabled() {
		instance.Use(mware.Logger())
	}

	if GetMiddlewareRecoverEnabled() {
		instance.Use(middleware.Recover())
	}

	if GetMiddlewareNewRelicEnabled() {
		instance.Use(nrecho.Middleware(newrelic.Application()))
	}

	if GetMiddlewarePrometheusEnabled() {
		instance.Use(prometheus.MetricsMiddleware())
	}

	if GetMiddlewareBodyDumpEnabled() {
		instance.Use(middleware.BodyDump(bodyDump))
	}

	if GetMiddlewareBodyLimitEnabled() {
		instance.Use(middleware.BodyLimit(GetMiddlewareBodyLimitSize()))
	}

	if GetMiddlewareCORSEnabled() {
		instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     GetMiddlewareCORSAllowOrigins(),
			AllowMethods:     GetMiddlewareCORSAllowMethods(),
			AllowHeaders:     GetMiddlewareCORSAllowHeaders(),
			AllowCredentials: GetMiddlewareCORSAllowCredentials(),
			ExposeHeaders:    GetMiddlewareCORSExposeHeaders(),
			MaxAge:           GetMiddlewareCORSMaxAge(),
		}))
	}

}

func setDefaultRouters(ctx context.Context, instance *echo.Echo) {

	logger := log.FromContext(ctx)

	statusRoute := GetStatusRoute()

	logger.Infof("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.GET(statusRoute, statusHandler.Get)

	healthRoute := GetHealthRoute()

	logger.Infof("configuring health router on %s", healthRoute)

	healthHandler := NewHealthHandler()
	instance.GET(healthRoute, healthHandler.Get)

	if GetMiddlewarePrometheusEnabled() {
		prometheusRoute := GetPrometheusRoute()
		logger.Infof("configuring prometheus metrics router on %s", prometheusRoute)
		instance.GET(prometheusRoute, echo.WrapHandler(promhttp.Handler()))
	}
}

func Serve(ctx context.Context) {
	l := log.FromContext(ctx)
	l.Infof("starting echo server. https://echo.labstack.com/")
	err := instance.Start(serverPort())
	if err != nil {
		instance.Logger.Fatalf(err.Error())
	}
}

func serverPort() string {
	return ":" + strconv.Itoa(GetPort())
}
