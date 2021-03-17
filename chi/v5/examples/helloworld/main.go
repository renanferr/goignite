package main

import (
	"context"
	"encoding/json"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gichihealth "github.com/b2wdigital/goignite/v2/chi/v5/ext/health"
	gichilogger "github.com/b2wdigital/goignite/v2/chi/v5/ext/logger"
	gichirealip "github.com/b2wdigital/goignite/v2/chi/v5/ext/realip"
	gichirecoverer "github.com/b2wdigital/goignite/v2/chi/v5/ext/recoverer"
	gichistatus "github.com/b2wdigital/goignite/v2/chi/v5/ext/status"
	gichitid "github.com/b2wdigital/goignite/v2/chi/v5/ext/tid"
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giinfo "github.com/b2wdigital/goignite/v2/info"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
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

	giinfo.AppName = "helloworld"

	srv := gichi.NewDefault(ctx,
		gichitid.Register,
		gichirecoverer.Register,
		gichirealip.Register,
		gichilogger.Register,
		gichistatus.Register,
		gichihealth.Register)

	srv.Mux().Get(c.App.Endpoint.Helloworld, Get(ctx))

	srv.Serve(ctx)
}
