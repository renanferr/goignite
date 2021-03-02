package cors

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if isEnabled() {
		instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     getAllowOrigins(),
			AllowMethods:     getAllowMethods(),
			AllowHeaders:     getAllowHeaders(),
			AllowCredentials: getAllowCredentials(),
			ExposeHeaders:    getExposeHeaders(),
			MaxAge:           getMaxAge(),
		}))
	}

	return nil
}
