package response

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UnprocessableEntityErrorResponse struct {
	ErrorResponse
	ValidationErrors []ValidationErrorResponse `json:"validationErrors,omitempty"`
}

func NewUnprocessableEntityErrorResponse(err validator.ValidationErrors) UnprocessableEntityErrorResponse {

	var fe validator.FieldError
	var verrs []ValidationErrorResponse

	for i := 0; i < len(err); i++ {

		fe = err[i].(validator.FieldError)

		verr := ValidationErrorResponse{
			FieldName: fe.Field(),
			Message:   "invalid value",
		}

		verrs = append(verrs, verr)
	}

	return UnprocessableEntityErrorResponse{
		ErrorResponse: ErrorResponse{
			HttpStatusCode: http.StatusUnprocessableEntity,
			Message:        "The server understands the content type of the request entity but was unable to process the contained instructions.",
		},
		ValidationErrors: verrs,
	}
}