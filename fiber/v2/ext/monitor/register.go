package gifibermonitor

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Register(ctx context.Context, app *fiber.App) error {
	if IsEnabled() {
		app.Use(monitor.New())
	}

	return nil
}
