package gichi

import (
	"context"
	"net/http"

	gihttp "github.com/b2wdigital/goignite/v2/http/v1/server"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-chi/chi/v5"
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

type Server struct {
	mux  *chi.Mux
	opts *gihttp.Options
}

func NewDefault(ctx context.Context, exts ...Ext) *Server {
	opt, err := gihttp.DefaultOptions()
	if err != nil {
		panic(err)
	}
	return New(ctx, opt, exts...)
}

func New(ctx context.Context, opts *gihttp.Options, exts ...Ext) *Server {

	mux := chi.NewRouter()

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
		mux.Use(middlewares...)
	}

	if len(handlers) > 0 {
		for _, h := range handlers {
			mux.Handle(h.Pattern, h.Handler)
		}
	}

	if len(routes) > 0 {
		for _, r := range routes {
			mux.MethodFunc(r.Method, r.Pattern, r.HandlerFunc)
		}
	}

	return &Server{mux: mux, opts: opts}
}

func (s *Server) Mux() *chi.Mux {
	return s.mux
}

func (s *Server) Serve(ctx context.Context) {

	logger := gilog.FromContext(ctx)

	httpServer := gihttp.New(s.mux, s.opts)

	logger.Infof("started chi http Server [%s]", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != nil {
		logger.Fatalf("cannot start chi http server", err.Error())
	}
}
