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

	logger.Infof("configuring status router on %s", statusRoute)

	app.Get(statusRoute, func(c *fiber.Ctx) error {
		return gifiber.JSON(c, http.StatusOK, response.NewResourceStatus(), nil)
	})

	return nil
}
