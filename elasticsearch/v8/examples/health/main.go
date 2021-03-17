package main

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/elasticsearch/v8"
	"github.com/b2wdigital/goignite/v2/elasticsearch/v8/ext/health"
	h "github.com/b2wdigital/goignite/v2/health"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

	integrator := health.NewDefaultIntegrator()

	_, err := elasticsearch.NewDefaultClient(context.Background(), integrator.Register)
	if err != nil {
		log.Panic(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
