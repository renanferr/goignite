package echo

import (
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/labstack/echo/v4"
)

func bodyDump(c echo.Context, reqBody []byte, resBody []byte) {
	logger := log.FromContext(c.Request().Context())
	logger.Info("request body --->")
	logger.Info(string(reqBody))
	logger.Info("response body -->")
	logger.Info(string(resBody))
}
