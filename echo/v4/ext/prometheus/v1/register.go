package giechoprometheus

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	prometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Register(ctx context.Context, instance *echo.Echo) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating echo with prometheus")

	instance.Use(prometheus.MetricsMiddleware())

	prometheusRoute := GetRoute()

	logger.Infof("configuring prometheus metrics router on %s", prometheusRoute)
	instance.GET(prometheusRoute, echo.WrapHandler(promhttp.Handler()))

	logger.Infof("prometheus integrated with echo with success")

	return nil
}
