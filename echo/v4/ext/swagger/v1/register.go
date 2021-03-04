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

	logger.Trace("integrating echo with swagger")

	swaggerRoute := GetRoute()
	logger.Infof("configuring swagger router on %s", swaggerRoute)

	instance.GET(swaggerRoute, echoSwagger.WrapHandler)

	logger.Debug("swagger integrated with echo with success")

	return nil
}
