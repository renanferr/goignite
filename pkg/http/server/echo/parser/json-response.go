package parser

import (
	"net/http"

	"github.com/jpfaria/goignite/pkg/errors"
	"github.com/jpfaria/goignite/pkg/http/server/model/response"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

func JSONResponse(c echo.Context, code int, i interface{}, err error) error {

	if err != nil {

		return JSONErrorResponse(c, err)

	}

	return c.JSONPretty(code, i, "  ")
}

func JSONErrorResponse(c echo.Context, err error) error {

	switch t := err.(type) {
	default:
		return c.JSONPretty(http.StatusInternalServerError, response.ErrorResponse{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()}, "  ")
	case *errors.NotFoundError:
		return c.JSONPretty(http.StatusNotFound, response.ErrorResponse{HttpStatusCode: http.StatusNotFound, Message: t.Error()}, "  ")
	case *errors.IllegalArgumentError:
		return c.JSONPretty(http.StatusBadRequest, response.ErrorResponse{HttpStatusCode: http.StatusBadRequest, Message: t.Error()}, "  ")
	case validator.ValidationErrors:

		var fe validator.FieldError
		var verrs []response.ValidationErrorResponse

		for i := 0; i < len(t); i++ {

			fe = t[i].(validator.FieldError)

			verr := response.ValidationErrorResponse{
				FieldName: fe.Field(),
				Message:   fe.Namespace(),
			}

			verrs = append(verrs, verr)
		}

		return c.JSONPretty(
			http.StatusUnprocessableEntity,
			response.UnprocessableEntityErrorResponse{
				ErrorResponse: response.ErrorResponse{
					HttpStatusCode: http.StatusUnprocessableEntity,
					Message:        "The server understands the content type of the request entity but was unable to process the contained instructions.",
				},
				ValidationErrors: verrs,
			},
			"  ",
		)
	}
}
