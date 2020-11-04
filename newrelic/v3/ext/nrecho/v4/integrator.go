package ginrecho

import (
	giecho "github.com/b2wdigital/goignite/echo/v4"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

type Integrator struct {
}

func Integrate() error {
	if !IsEnabled() {
		return nil
	}

	integrator := &Integrator{}
	return gieventbus.SubscribeOnce(giecho.TopicInstance, integrator.Integrate)
}

func (i *Integrator) Integrate(instance *echo.Echo) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating echo with newrelic")

	instance.Use(nrecho.Middleware(ginewrelic.Application()))

	if GetMiddlewareRequestIDEnabled() {
		logger.Debug("enabling newrelic requestID middleware")
		instance.Use(RequestIDMiddleware())
	}

	logger.Debug("echo integrated with newrelic with success")

	return nil
}
