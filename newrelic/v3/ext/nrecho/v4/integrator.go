package ginrecho

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

func Integrate(ctx context.Context, instance *echo.Echo) error {

	if isEnabled() {

		logger := gilog.FromContext(ctx)

		logger.Trace("integrating echo with newrelic")

		instance.Use(nrecho.Middleware(ginewrelic.Application()))

		if isEnabledRequestID() {
			logger.Debug("enabling newrelic requestID middleware")
			instance.Use(RequestIDMiddleware())
		}

		logger.Debug("echo integrated with newrelic with success")

	}

	return nil
}
