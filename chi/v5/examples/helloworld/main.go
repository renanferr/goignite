package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/b2wdigital/goignite/v2/chi/v5"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/health"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/logger"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/realip"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/recoverer"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/status"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/tid"
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/info"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"

func init() {
	config.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
}

type Config struct {
	App struct {
		Endpoint struct {
			Helloworld string
		}
	}
}

type Response struct {
	Message string
}

func Get(ctx context.Context) http.HandlerFunc {

	resp := Response{
		Message: "Hello World!!",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

func main() {

	config.Load()

	c := Config{}

	err := config.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	logrus.NewLogger()

	info.AppName = "helloworld"

	srv := chi.NewDefault(ctx,
		tid.Register,
		recoverer.Register,
		realip.Register,
		logger.Register,
		status.Register,
		health.Register)

	srv.Mux().Get(c.App.Endpoint.Helloworld, Get(ctx))

	srv.Serve(ctx)
}
