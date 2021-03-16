package gifiberprometheus

import (
	"context"

	"github.com/ansrivas/fiberprometheus/v2"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, instance *fiber.App) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("enabling prometheus middleware in fiber")

	prometheus := fiberprometheus.New("")
	instance.Use(prometheus.Middleware)

	logger.Debug("prometheus middleware successfully enabled in fiber")

	prometheusRoute := getRoute()

	logger.Tracef("configuring prometheus metric router on %s in fiber", prometheusRoute)

	prometheus.RegisterAt(instance, prometheusRoute)

	logger.Debugf("prometheus metric router configured on %s in fiber", prometheusRoute)

	return nil
}
