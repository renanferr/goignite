package ginrfiber

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/gofiber/fiber/v2"
)

func Integrate(ctx context.Context, instance *fiber.App) error {

	if !isEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating fiber with newrelic")

	instance.Use(middleware(ginewrelic.Application()))

	logger.Debug("fiber integrated with newrelic with success")

	return nil
}
