package gichi

import (
	"context"
	"net/http"

	gihttp "github.com/b2wdigital/goignite/v2/http/v1/server"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-chi/chi/v5"
)

var (
	instance *chi.Mux
)

type Config struct {
	Middlewares []func(http.Handler) http.Handler
	Handlers    []ConfigHandler
	Routes      []ConfigRouter
}

type ConfigHandler struct {
	Handler http.Handler
	Pattern string
}

type ConfigRouter struct {
	Method      string
	HandlerFunc http.HandlerFunc
	Pattern     string
}

type Ext func(context.Context) (*Config, error)

func New(ctx context.Context, exts ...Ext) *chi.Mux {

	instance = chi.NewRouter()

	var middlewares []func(http.Handler) http.Handler
	var handlers []ConfigHandler
	var routes []ConfigRouter

	for _, ext := range exts {

		var err error
		var config *Config

		if config, err = ext(ctx); err != nil {
			panic(err)
		}

		if config != nil {

			if len(config.Middlewares) > 0 {
				middlewares = append(middlewares, config.Middlewares...)
			}

			if len(config.Handlers) > 0 {
				handlers = append(handlers, config.Handlers...)
			}

			if len(config.Routes) > 0 {
				routes = append(routes, config.Routes...)
			}

		}
	}

	if len(middlewares) > 0 {
		instance.Use(middlewares...)
	}

	if len(handlers) > 0 {
		for _, h := range handlers {
			instance.Handle(h.Pattern, h.Handler)
		}
	}

	if len(routes) > 0 {
		for _, r := range routes {
			instance.MethodFunc(r.Method, r.Pattern, r.HandlerFunc)
		}
	}

	return instance
}

func Serve(ctx context.Context) {
	logger := gilog.FromContext(ctx)
	logger.Infof("started chi server [%s]", gihttp.GetServerAddress())
	if err := gihttp.New(instance).ListenAndServe(); err != nil {
		logger.Fatalf("cannot start chi server", err)
	}
}
