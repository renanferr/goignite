package giechoerrorhandler

import (
	"context"
	"fmt"
	"net/http"

	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
	"github.com/b2wdigital/goignite/v2/rest/response"
	"github.com/labstack/echo/v4"
)

func Register(ctx context.Context, instance *echo.Echo) error {
	if IsEnabled() {
		instance.HTTPErrorHandler = customHTTPErrorHandler
	}

	return nil
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var msg interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	} else {
		msg = err.Error()
	}

	resp := response.Error{HttpStatusCode: code, Message: fmt.Sprintf("%v", msg)}
	if err := giecho.JSON(c, code, resp, nil); err != nil {
		c.Logger().Error(err)
	}
}
