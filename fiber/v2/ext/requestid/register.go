package requestid

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Register(ctx context.Context, app *fiber.App) error {
	if isEnabled() {
		app.Use(requestid.New())
	}

	return nil
}
