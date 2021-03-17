package health

import (
	"context"

	"github.com/b2wdigital/goignite/rest/response"
	"github.com/b2wdigital/goignite/v2/fiber/v2"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	healthRoute := getRoute()

	logger.Tracef("configuring health router on %s in fiber", healthRoute)

	app.Get(healthRoute, func(c *fiber.Ctx) error {

		ctx, cancel := context.WithCancel(c.Context())
		defer cancel()

		resp, httpCode := response.NewHealth(ctx)

		return fiber.JSON(c, httpCode, resp, nil)
	})

	logger.Debugf("health router configured on %s in fiber", healthRoute)

	return nil
}
