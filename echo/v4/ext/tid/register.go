package tid

import (
	"context"

	"github.com/b2wdigital/goignite/info"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if isEnabled() {
		instance.Use(tid())
	}

	return nil
}

// tid returns a middleware that tid HTTP requests.
func tid() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			ctx := c.Request().Context()

			tid := ctx.Value("x-tid")
			if tid == "" {
				tid = info.AppName + "-" + uuid.NewV4().String()
			}

			c.Request().WithContext(context.WithValue(ctx, "x-tid", tid))

			return nil
		}
	}
}
