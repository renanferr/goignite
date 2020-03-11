package main

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/info"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/b2wdigital/goignite/pkg/transport/client/http/resty/v2"
	e "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const Endpoint = "app.endpoint.google"

func init() {
	config.Add(Endpoint, "/google", "google endpoint")
}

type Config struct {
	App struct {
		Endpoint struct {
			Google string
		}
	}
}

type Response struct {
	Message string
}

type Handler struct {
	client *resty.Client
}

func NewHandler(client *resty.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) Get(c e.Context) (err error) {

	l := log.FromContext(c.Request().Context())

	request := h.client.R().EnableTrace()

	_, err = request.Get("http://google.com")
	if err != nil {
		l.Fatalf(err)
	}

	resp := Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Errorf(err.Error())
	}

	return v4.JSONResponse(c, http.StatusOK, resp, err)
}

func main() {

	err := config.Load()
	if err != nil {
		panic(err)
	}

	c := Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	log.NewLogger(v1.NewLogger())

	info.AppName = "google"

	instance := v4.Start(ctx)

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	//instance.AddErrorAdvice(customErrors.InvalidPayload, 400)

	o := resty.OptionsBuilder.
		Host("http://www.google.com").
		Health(
			resty.OptionsHealthBuilder.
				Enabled(true).
				Required(false).
				Description("google dependency").
				Endpoint("/").
				Build()).
		Build()

	client := resty.NewClient(ctx, &o)

	handler := NewHandler(client)
	instance.GET(c.App.Endpoint.Google, handler.Get)

	v4.Serve(ctx)
}
