package giecho

import (
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/labstack/echo/v4"
)

func bodyDump(c echo.Context, reqBody []byte, resBody []byte) {
	logger := gilog.FromContext(c.Request().Context())
	logger.Info("request body --->")
	logger.Info(string(reqBody))
	logger.Info("response body -->")
	logger.Info(string(resBody))
}
