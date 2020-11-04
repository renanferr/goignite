package main

import (
	"context"
	"encoding/json"
	"net/http"

	gichi "github.com/b2wdigital/goignite/chi/v4"
	giconfig "github.com/b2wdigital/goignite/config"
	gihttp "github.com/b2wdigital/goignite/http/v1/server"
	"github.com/b2wdigital/goignite/info"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	"github.com/prometheus/common/log"
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

	instance := gichi.NewMux(ctx)

	instance.Get(c.App.Endpoint.Helloworld, Get(ctx))

	log.Infof("starting chi server.")
	err = gihttp.NewServer(instance).ListenAndServe()
	log.Fatalf("cannot start chi server", err)

}
