package bodylimit

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Middleware(ctx context.Context, instance *echo.Echo) error {
	if isEnabled() {
		instance.Use(middleware.BodyLimit(getSize()))
	}

	return nil
}
