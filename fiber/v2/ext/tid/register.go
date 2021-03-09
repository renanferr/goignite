package gifibertid

import (
	"context"

	"github.com/b2wdigital/goignite/v2/info"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, app *fiber.App) error {
	if !IsEnabled() {
		return nil
	}
	app.Use(tidMiddleware())
	return nil
}

func tidMiddleware() fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) error {

		tid := c.Get("X-TID", info.AppName+"-"+uuid.NewV4().String())
		c.Context().SetUserValue("x-tid", tid)
		c.Append("X-TID", tid)

		// Continue stack
		return c.Next()
	}
}
