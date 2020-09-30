package giecho

import (
	"context"
	"strconv"

	mware "github.com/b2wdigital/goignite/echo/v4/middleware"
	gilog "github.com/b2wdigital/goignite/log"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	prometheus "github.com/globocom/echo-prometheus"
	echopprof "github.com/hiko1129/echo-pprof"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	instance *echo.Echo
)

func Start(ctx context.Context) *echo.Echo {

	instance = echo.New()

	instance.HideBanner = GetHideBanner()
	instance.Logger = Wrap(gilog.GetLogger())

	setDefaultMiddlewares(instance)
	setDefaultRouters(ctx, instance)

	return instance
}

func setDefaultMiddlewares(instance *echo.Echo) {

	if GetMiddlewareRequestIDEnabled() {
		instance.Use(middleware.RequestID())
	}

	if GetMiddlewareLogEnabled() {
		instance.Use(mware.Logger())
	}

	if GetMiddlewareRecoverEnabled() {
		instance.Use(middleware.Recover())
	}

	if GetMiddlewareNewRelicEnabled() {
		instance.Use(nrecho.Middleware(ginewrelic.Application()))
		instance.Use(mware.NewRelicAddonMiddleware())
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

	logger := gilog.FromContext(ctx)

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

	if GetMiddlewareSwaggerEnabled() {
		swaggerRoute := GetSwaggerRoute() + "/*"
		logger.Infof("configuring swagger router on %s", swaggerRoute)
		instance.GET(swaggerRoute, echoSwagger.WrapHandler)
	}

	if GetPProfEnabled() {
		echopprof.Wrap(instance)
	}
}

func Serve(ctx context.Context) {
	l := gilog.FromContext(ctx)
	l.Infof("starting echo server. https://echo.labstack.com/")
	err := instance.Start(serverPort())
	if err != nil {
		instance.Logger.Fatalf(err.Error())
	}
}

func serverPort() string {
	return ":" + strconv.Itoa(GetPort())
}
