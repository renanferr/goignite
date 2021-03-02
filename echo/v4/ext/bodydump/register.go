package bodydump

import (
	"context"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if isEnabled() {
		instance.Use(middleware.BodyDump(bodyDump))
	}

	return nil
}

func bodyDump(c echo.Context, reqBody []byte, resBody []byte) {
	logger := gilog.FromContext(c.Request().Context())
	logger.Info("request body --->")
	logger.Info(string(reqBody))
	logger.Info("response body -->")
	logger.Info(string(resBody))
}
