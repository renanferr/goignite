package validator

import (
	"reflect"
	"strings"

	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/br"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fr"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	fr_translations "github.com/go-playground/validator/v10/translations/fr"
	pt_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var translator ut.Translator

// Default locale for get the specified translator
var defaultLocale = "pt_BR"

const (
	REQUIRED = "REQUIRED"
	INVALID  = "INVALID"
)

// The options is a optional parameter, for default
// locale is pt_BR, but we can change using the SetLocale
func NewTranslator(validate *validator.Validate, options ...Options) {

	_options := Options{Locale: defaultLocale}

	for _, opt := range options {
		if opt.Locale != "" {
			_options.Locale = opt.Locale
		}

		_options.Schema = opt.Schema
		break
	}

	fallback := en.New()
	locales := supportedLocales()

	translator, _ = ut.New(fallback, locales...).GetTranslator(_options.Locale)

	switch _options.Locale {
	case "pt_BR":
		if err := pt_translations.RegisterDefaultTranslations(validate, translator); err != nil {
			log.Fatal(err.Error())
		}
	case "fr":
		if err := fr_translations.RegisterDefaultTranslations(validate, translator); err != nil {
			log.Fatal(err.Error())
		}
	default:
		if err := en_translations.RegisterDefaultTranslations(validate, translator); err != nil {
			log.Fatal(err.Error())
		}
	}

	for tag, message := range _options.Schema {
		_tag := tag
		fn := func(ut ut.Translator, fe validator.FieldError) string {
			if strings.ToUpper(_tag) == REQUIRED {
				fieldName, _ := ut.T(_tag, strings.ToLower(fe.Field()))
				return fieldName
			}

			// add more cases here if necessary
			switch value := fe.Value().(type) {
			case string:
				t, _ := ut.T(_tag, strings.ToLower(value))
				return t

			default:
				t, _ := ut.T(_tag, value.(string))
				return t
			}
		}

		validate.RegisterTranslation(tag, translator, func(ut ut.Translator) error {
			return ut.Add(tag, message, true)
		}, fn)
	}

}

func Translator() ut.Translator {
	return translator
}

func RegisterTagNameFunc(validate *validator.Validate) {

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// At future this function should be allows add/append a locales.Translator
func supportedLocales() []locales.Translator {

	return []locales.Translator{
		en.New(),
		br.New(),
		fr.New(),
	}
}
