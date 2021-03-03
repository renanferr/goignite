package gzip

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if IsEnabled() {
		instance.Use(middleware.Gzip())
	}

	return nil
}
