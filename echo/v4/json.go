package giecho

import (
	"net/http"

	gierrors "github.com/b2wdigital/goignite/v2/errors"
	girestresponse "github.com/b2wdigital/goignite/v2/rest/response"
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

	return json(c, code, i)
}

func JSONError(c echo.Context, err error) error {

	if gierrors.IsNotFound(err) {
		return json(c,
			http.StatusNotFound,
			girestresponse.Error{HttpStatusCode: http.StatusNotFound, Message: err.Error()})
	} else if gierrors.IsNotValid(err) || gierrors.IsBadRequest(err) {
		return json(c,
			http.StatusBadRequest,
			girestresponse.Error{HttpStatusCode: http.StatusBadRequest, Message: err.Error()})
	} else if gierrors.IsServiceUnavailable(err) {
		return json(c,
			http.StatusServiceUnavailable,
			girestresponse.Error{HttpStatusCode: http.StatusServiceUnavailable, Message: err.Error()})
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			return json(c,
				http.StatusUnprocessableEntity,
				girestresponse.NewUnprocessableEntity(t))
		default:
			return json(c,
				http.StatusInternalServerError,
				girestresponse.Error{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()})
		}
	}

}

func json(c echo.Context, code int, response interface{}) error {

	if GetJSONPrettyEnabled() {
		return c.JSONPretty(code, response, "  ")
	}

	return c.JSON(code, response)
}
