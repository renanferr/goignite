package recover

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Middleware(ctx context.Context, app *fiber.App) error {
	if isEnabled() {
		app.Use(recover.New())
	}

	return nil
}
