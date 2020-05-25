package giecho

import (
	"net/http"

	"github.com/b2wdigital/goignite/errors"
	"github.com/b2wdigital/goignite/pkg/rest/response"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func JSON(c echo.Context, code int, i interface{}, err error) error {

	if err != nil {

		return JSONError(c, err)

	}

	if i == nil {
		return c.NoContent(code)
	}

	return c.JSONPretty(code, i, "  ")
}

func JSONError(c echo.Context, err error) error {

	if errors.IsNotFound(err) {
		return c.JSONPretty(
			http.StatusNotFound,
			response.Error{HttpStatusCode: http.StatusNotFound, Message: err.Error()},
			"  ")
	} else if errors.IsNotValid(err) || errors.IsBadRequest(err) {
		return c.JSONPretty(
			http.StatusBadRequest,
			response.Error{HttpStatusCode: http.StatusBadRequest, Message: err.Error()},
			"  ")
	} else {

		switch t := err.(type) {
		default:
			return c.JSONPretty(
				http.StatusInternalServerError,
				response.Error{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()},
				"  ")
		case validator.ValidationErrors:
			return c.JSONPretty(
				http.StatusUnprocessableEntity,
				response.NewUnprocessableEntity(t),
				"  ")
		}
	}

}