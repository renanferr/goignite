package gifiber

import (
	"net/http"

	gierrors "github.com/b2wdigital/goignite/v2/errors"
	girestresponse "github.com/b2wdigital/goignite/v2/rest/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func JSON(c *fiber.Ctx, code int, i interface{}, err error) error {

	if err != nil {
		return JSONError(c, err)
	}

	if i == nil {
		c.Status(http.StatusNoContent)
		return nil
	}

	return c.Status(code).JSON(i)
}

func JSONError(c *fiber.Ctx, err error) error {

	if gierrors.IsNotFound(err) {
		return c.Status(http.StatusNotFound).JSON(
			girestresponse.Error{HttpStatusCode: http.StatusNotFound, Message: err.Error()})
	} else if gierrors.IsNotValid(err) || gierrors.IsBadRequest(err) {
		return c.Status(http.StatusBadRequest).JSON(
			girestresponse.Error{HttpStatusCode: http.StatusBadRequest, Message: err.Error()})
	} else if gierrors.IsServiceUnavailable(err) {
		return c.Status(http.StatusServiceUnavailable).JSON(
			girestresponse.Error{HttpStatusCode: http.StatusServiceUnavailable, Message: err.Error()})
	} else {

		switch t := err.(type) {
		case validator.ValidationErrors:
			return c.Status(http.StatusUnprocessableEntity).JSON(
				girestresponse.NewUnprocessableEntity(t))
		default:
			return c.Status(http.StatusInternalServerError).JSON(
				girestresponse.Error{HttpStatusCode: http.StatusInternalServerError, Message: t.Error()})
		}
	}

}
