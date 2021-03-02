package logger

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Middleware(ctx context.Context, app *fiber.App) error {
	if isEnabled() {
		app.Use(logger.New())
	}

	return nil
}
