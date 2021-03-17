package datadog

import (
	"context"

	"github.com/b2wdigital/goignite/v2/datadog/v1"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/labstack/echo/v4"
	ddecho "gopkg.in/DataDog/dd-trace-go.v1/contrib/labstack/echo.v4"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if !IsEnabled() || !datadog.IsEnabled() {
		return nil
	}

	logger := log.FromContext(ctx)

	logger.Trace("enabling datadog middleware in echo")

	instance.Use(ddecho.Middleware())

	logger.Debug("datadog middleware successfully enabled in echo")

	return nil
}
