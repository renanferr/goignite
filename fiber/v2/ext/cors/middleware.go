package cors

import (
	"context"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Middleware(ctx context.Context, app *fiber.App) error {

	if !isEnabled() {
		return nil
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Join(getAllowOrigins(), ","),
		AllowMethods:     strings.Join(getAllowMethods(), ","),
		AllowHeaders:     strings.Join(getAllowHeaders(), ","),
		AllowCredentials: getAllowCredentials(),
		ExposeHeaders:    strings.Join(getExposeHeaders(), ","),
		MaxAge:           getMaxAge(),
	}))

	return nil
}
