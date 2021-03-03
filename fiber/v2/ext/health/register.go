package health

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/b2wdigital/goignite/rest/response"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	healthRoute := getRoute()

	logger.Infof("configuring health router on %s", healthRoute)

	app.Get(healthRoute, func(c *fiber.Ctx) error {

		ctx, cancel := context.WithCancel(c.Context())
		defer cancel()

		resp, httpCode := response.NewHealth(ctx)

		return c.Status(httpCode).JSON(resp)
	})

	return nil
}
