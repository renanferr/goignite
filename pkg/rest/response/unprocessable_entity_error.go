package response

import (
	"net/http"
	"regexp"
	"strings"

	v "github.com/b2wdigital/goignite/pkg/validator"
	"github.com/go-playground/validator/v10"
)

type UnprocessableEntityError struct {
	Error
	ValidationErrors []ValidationError `json:"validationErrors,omitempty"`
}

func NewUnprocessableEntity(err validator.ValidationErrors) UnprocessableEntityError {

	var fe validator.FieldError
	var verrs []ValidationError

	for i := 0; i < len(err); i++ {

		fe = err[i].(validator.FieldError)

		restrictionType := v.REQUIRED
		if strings.ToUpper(fe.Tag()) != v.REQUIRED {
			restrictionType = v.INVALID
		}

		verr := ValidationError{
			FieldName:       transformFieldName(fe.Namespace()),
			RestrictionType: restrictionType,
			Message:         transformMessage(fe, err.Error()),
		}

		verrs = append(verrs, verr)
	}

	return UnprocessableEntityError{
		Error: Error{
			HttpStatusCode: http.StatusUnprocessableEntity,
			Message:        "The server understands the content type of the request entity but was unable to process the contained instructions.",
		},
		ValidationErrors: verrs,
	}
}

func transformMessage(fe validator.FieldError, errorMessage string) string {

	instance := v.Translator()
	if instance == nil {
		return "invalid value"
	}

	message := fe.Translate(instance)

	// When don't exists message registered,
	// override default

	if strings.Contains(errorMessage, message) {
		return "invalid value"
	}

	return message
}

func transformFieldName(fieldName string) string {

	re := regexp.MustCompile("^(.*?)\\.(.*)$")
	return re.ReplaceAllString(fieldName, "$2")
}
