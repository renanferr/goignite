package monitor

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Middleware(ctx context.Context, app *fiber.App) error {
	if isEnabled() {
		app.Use(monitor.New())
	}

	return nil
}
