package ginrfiber

import (
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gifiber "github.com/b2wdigital/goignite/fiber/v2"
	gilog "github.com/b2wdigital/goignite/log"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/gofiber/fiber/v2"
)

type Integrator struct {
}

func Integrate() error {
	if !IsEnabled() {
		return nil
	}

	integrator := &Integrator{}
	return gieventbus.SubscribeOnce(gifiber.TopicApp, integrator.Integrate)
}

func (i *Integrator) Integrate(instance *fiber.App) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating fiber with newrelic")

	instance.Use(middleware(ginewrelic.Application()))

	logger.Debug("fiber integrated with newrelic with success")

	return nil
}
