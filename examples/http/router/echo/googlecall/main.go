package main

import (
	"context"
	"log"
	"net/http"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/http/client/resty/model"
	r "github.com/b2wdigital/goignite/pkg/http/client/resty/v2"
	"github.com/b2wdigital/goignite/pkg/http/router/echo"
	"github.com/b2wdigital/goignite/pkg/http/router/echo/parser"
	"github.com/b2wdigital/goignite/pkg/info"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"github.com/go-resty/resty/v2"
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

	l := logrus.FromContext(c.Request().Context())

	request := h.client.R().EnableTrace()

	_, err = request.Get("http://google.com")
	if err != nil {
		l.Fatal(err)
	}

	resp := Response{
		Message: "Hello Google!!",
	}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Error(err)
	}

	return parser.JSONResponse(c, http.StatusOK, resp, err)
}

func main() {

	err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	c := Config{}

	err = config.Unmarshal(&c)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	logrus.Start()

	info.AppName = "google"

	instance := echo.Start(ctx)

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	o := model.OptionsBuilder.
		Host("http://www.googeeele.com").
		Health(
			model.OptionsHealthBuilder.
				Enabled(true).
				Required(true).
				Description("google dependency").
				Endpoint("/").
				Build()).
		Build()

	client := r.NewClient(ctx, &o)

	handler := NewHandler(client)
	instance.GET(c.App.Endpoint.Google, handler.Get)

	echo.Serve(ctx)
}
