package main

import (
	"context"
	"encoding/json"
	"log"

    "github.com/go-playground/validator/v10"
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/http/server/echo"
	"github.com/b2wdigital/goignite/pkg/http/server/echo/parser"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/cloudevents/sdk-go"
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

{{ range .RequestMaps }}
func {{.Handler.Alias}}{{.Handler.Func}}{{.Method}}CloudEventsWrapper(ctx context.Context, e cloudevents.Event, resp *cloudevents.EventResponse) error {
    h := {{.Handler.Alias}}.{{.Handler.Func}}

    return h(ctx, e, resp)
}

func {{.Handler.Alias}}{{.Handler.Func}}{{.Method}}EchoWrapper(c e.Context) error {

    var err error

    h := {{.Handler.Alias}}{{.Handler.Func}}{{.Method}}CloudEventsWrapper

    request := cloudevents.Event{
        Context: cloudevents.EventContextV03{
            Source: *cloudevents.ParseURLRef("/mod3"),
            Type:   "samples.http.mod3",
        }.AsV03(),
    }

    {{- if or (eq .Method "POST") (eq .Method "PUT") (eq .Method "PATCH") }}

    st := new({{.Body.Alias}}.{{.Body.Struct}})

    if err = c.Bind(&st); err != nil {
        return parser.JSONErrorResponse(c, err)
    }

    if err = c.Validate(st); err != nil {
        return parser.JSONErrorResponse(c, err)
    }

    var str []byte

    str, err = json.Marshal(st)
    if err != nil {
        return parser.JSONErrorResponse(c, err)
    }
    request.Data = str

    {{- end }}

    resp := cloudevents.EventResponse{
        Context: cloudevents.EventContextV03{
            Source: *cloudevents.ParseURLRef("/mod3"),
            Type:   "samples.http.mod3",
        }.AsV03(),
    }

    err = h(c.Request().Context(), request, &resp)

	return parser.JSONResponse(c, {{.HttpCode}}, resp.Event.Data, err)
}
{{ end }}

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
	instance.{{.Method}}("{{.Endpoint}}", {{.Handler.Alias}}{{.Handler.Func}}{{.Method}}EchoWrapper)
	{{- end}}

	echo.Serve(ctx)
}