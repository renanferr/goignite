package newrelic

import (
	"context"

	"github.com/b2wdigital/goignite/v2/datadog/v1"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"
)

func Register(ctx context.Context, instance *fiber.App) error {

	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)
	logger.Trace("enabling datadog middleware in fiber")

	instance.Use(fibertrace.Middleware())

	logger.Debug("datadog middleware successfully enabled in fiber")

	return nil
}
