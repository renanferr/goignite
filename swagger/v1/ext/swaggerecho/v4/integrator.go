package giswaggerecho

import (
	giecho "github.com/b2wdigital/goignite/echo/v4"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/labstack/echo/v4"
	// echoSwagger "github.com/swaggo/echo-swagger"
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

	logger.Trace("integrating echo with swagger")

	swaggerRoute := GetRoute() + "/*"
	logger.Infof("configuring swagger router on %s", swaggerRoute)
	// instance.GET(swaggerRoute, echoSwagger.WrapHandler)

	logger.Debug("swagger integrated with echo with success")

	return nil
}
