package pprof

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func Register(ctx context.Context, app *fiber.App) error {
	if IsEnabled() {
		app.Use(pprof.New())
	}

	return nil
}
