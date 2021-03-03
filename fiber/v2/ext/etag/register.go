package etag

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func Register(ctx context.Context, app *fiber.App) error {
	if IsEnabled() {
		app.Use(etag.New())
	}

	return nil
}
