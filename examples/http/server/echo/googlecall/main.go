package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/jpfaria/goignite/internal/pkg/info"
	"github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/http/client/resty/model"
	r "github.com/jpfaria/goignite/pkg/http/client/resty/v2"
	"github.com/jpfaria/goignite/pkg/http/server/echo"
	"github.com/jpfaria/goignite/pkg/http/server/echo/parser"
	"github.com/jpfaria/goignite/pkg/logging/logrus"
	e "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	if err!= nil {
		l.Fatal(err)
	}

	resp := Response{}

	err = config.Unmarshal(&resp)
	if err != nil {
		l.Error(err)
	}

	return parser.JSONResponse(c, http.StatusOK, resp, err)
}

func main() {

	err := config.Parse()
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

	instance := echo.Start()

	instance.Use(middleware.Gzip())
	instance.Use(middleware.CORS())
	instance.Use(middleware.RequestID())

	o := model.OptionsBuilder.
		Host("http://www.google.com").
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

	echo.Serve()
}
