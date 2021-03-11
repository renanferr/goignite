package giechoswagger

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Register(ctx context.Context, instance *echo.Echo) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	swaggerRoute := GetRoute()

	logger.Tracef("configuring swagger router on %s in echo", swaggerRoute)

	instance.GET(swaggerRoute, echoSwagger.WrapHandler)

	logger.Debugf("swagger router configured on %s in echo", swaggerRoute)

	return nil
}
