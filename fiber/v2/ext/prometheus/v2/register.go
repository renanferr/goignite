package prometheus

import (
	"context"

	"github.com/ansrivas/fiberprometheus/v2"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, instance *fiber.App) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating fiber with prometheus")

	prometheusRoute := getRoute()

	logger.Infof("configuring prometheus metrics router on %s", prometheusRoute)

	prometheus := fiberprometheus.New("")
	prometheus.RegisterAt(instance, prometheusRoute)

	instance.Use(prometheus.Middleware)

	logger.Debug("prometheus integrated with fiber with success")

	return nil
}
