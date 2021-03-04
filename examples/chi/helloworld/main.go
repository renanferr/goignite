package main

import (
	"context"
	"encoding/json"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/health"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/logger"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/realip"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/recoverer"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/status"
	"github.com/b2wdigital/goignite/v2/chi/v5/ext/tid"
	giconfig "github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/info"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
)

const HelloWorldEndpoint = "app.endpoint.helloworld"

func init() {
	giconfig.Add(HelloWorldEndpoint, "/hello-world", "helloworld endpoint")
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

	giconfig.Load()

	c := Config{}

	err := giconfig.Unmarshal(&c)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	gilogrus.NewLogger()

	info.AppName = "helloworld"

	instance := gichi.New(ctx,
		tid.Register,
		recoverer.Register,
		realip.Register,
		logger.Register,
		status.Register,
		health.Register)

	instance.Get(c.App.Endpoint.Helloworld, Get(ctx))

	gichi.Serve(ctx)
}
