package gifiberlogger

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}

	l := gilog.FromContext(ctx)
	l.Trace("enabling logger middleware in fiber")

	app.Use(logger.New(logger.Config{
		Output: gilog.GetLogger().Output(),
	}))

	l.Debug("logger middleware successfully enabled in fiber")

	return nil
}
