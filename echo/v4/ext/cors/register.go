package giechocors

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if IsEnabled() {
		instance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     GetAllowOrigins(),
			AllowMethods:     GetAllowMethods(),
			AllowHeaders:     GetAllowHeaders(),
			AllowCredentials: GetAllowCredentials(),
			ExposeHeaders:    GetExposeHeaders(),
			MaxAge:           GetMaxAge(),
		}))
	}

	return nil
}
