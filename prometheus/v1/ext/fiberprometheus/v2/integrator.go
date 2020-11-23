package giprometheusfiber

import (
	"github.com/ansrivas/fiberprometheus/v2"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gifiber "github.com/b2wdigital/goignite/fiber/v2"
	gilog "github.com/b2wdigital/goignite/log"
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

	logger.Trace("integrating fiber with prometheus")

	prometheusRoute := GetRoute()

	logger.Infof("configuring prometheus metrics router on %s", prometheusRoute)

	prometheus := fiberprometheus.New("graphql-go-server")
	prometheus.RegisterAt(instance, prometheusRoute)

	instance.Use(prometheus.Middleware)

	logger.Debug("prometheus integrated with fiber with success")

	return nil
}
