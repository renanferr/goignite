package gipromecho

import (
	giecho "github.com/b2wdigital/goignite/echo/v4"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	prometheus "github.com/globocom/echo-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Integrator struct {
}

func Integrate() error {
	integrator := &Integrator{}
	return gieventbus.SubscribeOnce(giecho.TopicInstance, integrator.Integrate)
}

func (i *Integrator) Integrate(instance *echo.Echo) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating echo with prometheus")

	if IsEnabled() {

		instance.Use(prometheus.MetricsMiddleware())

		prometheusRoute := GetRoute()

		logger.Infof("configuring prometheus metrics router on %s", prometheusRoute)
		instance.GET(prometheusRoute, echo.WrapHandler(promhttp.Handler()))
	}

	logger.Debug("prometheus integrated with echo with success")

	return nil
}
