package middleware

import (
	"fmt"

	"github.com/b2wdigital/goignite/pkg/transport/client/newrelic/v3"
	"github.com/labstack/echo/v4"
)

const (
	// NEWRELIC_TXN defines the context key used to save newrelic transaction
	NEWRELIC_TXN = "newrelic-txn"
)

func NewRelic() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			transactionName := fmt.Sprintf("%s [%s]", c.Path(), c.Request().Method)
			txn := newrelic.Application().StartTransaction(transactionName)
			defer txn.End()

			c.Set(NEWRELIC_TXN, txn)

			err := next(c)

			if err != nil {
				txn.NoticeError(err)
			}

			txn.End()

			return err
		}
	}
}
