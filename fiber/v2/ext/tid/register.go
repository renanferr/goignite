package tid

import (
	"context"

	"github.com/b2wdigital/goignite/info"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, app *fiber.App) error {

	if !isEnabled() {
		return nil
	}

	app.Use(middleware())

	return nil
}

func middleware() fiber.Handler {

	// Return new handler
	return func(c *fiber.Ctx) error {

		tid := c.Get("X-TID", info.AppName+"-"+uuid.NewV4().String())

		c.Context().SetUserValue("tid", tid)
		c.Append("X-TID", tid)

		// Continue stack
		return c.Next()
	}
}
