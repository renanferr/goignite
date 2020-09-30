package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicAddonMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			ctx := c.Request().Context()
			txn := newrelic.FromContext(ctx)
			reqId := c.Request().Header.Get(echo.HeaderXRequestID)
			if reqId == "" {
				reqId = c.Response().Header().Get(echo.HeaderXRequestID)
			}

			txn.AddAttribute("request.id", reqId)
			return next(c)
		}
	}
}
