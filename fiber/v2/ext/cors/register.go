package gifibercors

import (
	"context"
	"strings"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Register(ctx context.Context, app *fiber.App) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("enabling cors middleware in fiber")

	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(getAllowOrigins(), ","),
		AllowMethods:     strings.Join(getAllowMethods(), ","),
		AllowHeaders:     strings.Join(getAllowHeaders(), ","),
		AllowCredentials: getAllowCredentials(),
		ExposeHeaders:    strings.Join(getExposeHeaders(), ","),
		MaxAge:           getMaxAge(),
	}))

	logger.Debug("cors middleware successfully enabled in fiber")

	return nil
}
