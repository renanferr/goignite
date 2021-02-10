package ginrfiber

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func middleware(app *newrelic.Application) fiber.Handler {

	if app == nil {
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	return func(c *fiber.Ctx) error {
		transactionPattern := fmt.Sprintf("%s - %s ", c.Method(), string(c.Request().URI().Path()))
		txn := app.StartTransaction(transactionPattern)
		defer txn.End()

		// TODO criar whitelist de headers
		c.Request().Header.VisitAll(func(key, value []byte) {
			txn.AddAttribute(strings.ToLower(string(key)), string(value))
		})

		wr := setNewRelicWebRequest(c)
		txn.SetWebRequest(wr)

		c.Locals("txn", txn)

		return c.Next()
	}
}

func setNewRelicWebRequest(c *fiber.Ctx) newrelic.WebRequest {
	header := http.Header{}

	c.Request().Header.VisitAll(func(key, value []byte) {
		header.Add(string(key), string(value))
	})

	URL := fmt.Sprintf("%s%s", c.BaseURL(), c.Path())
	parsedURL, _ := url.Parse(URL)

	wr := newrelic.WebRequest{
		Header:    header,
		URL:       parsedURL,
		Method:    c.Method(),
		Transport: newrelic.TransportType(c.Protocol()),
		Host:      string(c.Request().Host()),
	}

	return wr
}
