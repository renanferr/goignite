package status

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/b2wdigital/goignite/rest/response"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !isEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	statusRoute := getRoute()

	logger.Infof("configuring status router on %s", statusRoute)

	app.Get(statusRoute, func(c *fiber.Ctx) error {
		return c.JSON(response.NewResourceStatus())
	})

	return nil
}
