package gifiberlogger

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Register(ctx context.Context, app *fiber.App) error {
	if IsEnabled() {
		app.Use(logger.New())
	}

	return nil
}
