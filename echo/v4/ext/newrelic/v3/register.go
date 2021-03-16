package giechonewrelic

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	ginewrelic "github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

func Register(ctx context.Context, instance *echo.Echo) error {

	if !IsEnabled() || !ginewrelic.IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("enabling newrelic middleware in echo")

	instance.Use(nrecho.Middleware(ginewrelic.Application()))

	logger.Debug("newrelic middleware successfully enabled in echo")

	if IsEnabledRequestID() {

		logger.Trace("enabling requestID newrelic middleware in echo")

		instance.Use(requestIDMiddleware())

		logger.Debug("requestID newrelic middleware successfully enabled in echo")
	}

	return nil
}

func requestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ctx := c.Request().Context()
			txn := ginewrelic.FromContext(ctx)
			reqId := c.Request().Header.Get(echo.HeaderXRequestID)
			if reqId == "" {
				reqId = c.Response().Header().Get(echo.HeaderXRequestID)
			}

			txn.AddAttribute("request.id", reqId)
			return next(c)
		}
	}
}
