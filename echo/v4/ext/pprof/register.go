package pprof

import (
	"context"

	gilog "github.com/b2wdigital/goignite/v2/log"
	echopprof "github.com/hiko1129/echo-pprof"
	"github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *echo.Echo) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating echo with pprof")

	echopprof.Wrap(instance)

	logger.Debug("pprof integrated with echo with success")

	return nil
}
