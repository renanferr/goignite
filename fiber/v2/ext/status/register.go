package gifiberstatus

import (
	"context"
	"net/http"

	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/rest/response"
	"github.com/gofiber/fiber/v2"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	statusRoute := getRoute()

	logger.Tracef("configuring status router on %s in fiber", statusRoute)

	app.Get(statusRoute, func(c *fiber.Ctx) error {
		return gifiber.JSON(c, http.StatusOK, response.NewResourceStatus(), nil)
	})

	logger.Debugf("status router configured on %s in fiber", statusRoute)

	return nil
}
