package main

import (
	"context"
	"log"
	"net/http"

    "github.com/go-playground/validator/v10"
	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/http/server/echo"
	"github.com/jpfaria/goignite/pkg/http/server/echo/parser"
	"github.com/jpfaria/goignite/pkg/log/logrus"
	e "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	{{- range .Packages}}
	{{ .Alias}} "{{ .URI}}"
	{{- end}}
)

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

{{ range .RequestMaps}}
func {{.Handler.Alias}}{{.Handler.Func}}{{.Method}}(c e.Context) error {

    h := {{.Handler.Alias}}.{{.Handler.Func}}

    {{- if or (eq .Method "POST") (eq .Method "PUT") (eq .Method "PATCH") }}

    st := new({{.Body.Alias}}.{{.Body.Struct}})

    if err := c.Bind(&st); err != nil {
        return parser.JSONErrorResponse(c, err)
    }

    if err := c.Validate(st); err != nil {
        return parser.JSONErrorResponse(c, err)
    }

    response, err := h(c.Request(), st)

    {{- else }}

    response, err := h(c.Request())

    {{- end}}

    {{- if (eq .Method "GET") }}
    statusCode := http.StatusOK
    {{- else if (eq .Method "POST") }}
    statusCode := http.StatusCreated
    {{- else if (eq .Method "PUT") }}
    statusCode := http.StatusNoContent
    {{- else if (eq .Method "PATCH") }}
    statusCode := http.StatusNoContent
    {{- else if (eq .Method "DELETE") }}
    statusCode := http.StatusNoContent
    {{- end}}

	return parser.JSONResponse(c, statusCode, response, err)
}
{{- end}}

func main() {

	err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	logrus.Start()

	instance := echo.Start(ctx)

	instance.Validator = &CustomValidator{validator: validator.New()}

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

    {{ range .RequestMaps}}
	instance.{{.Method}}("{{.Endpoint}}", {{.Handler.Alias}}{{.Handler.Func}}{{.Method}})
	{{- end}}

	echo.Serve(ctx)
}