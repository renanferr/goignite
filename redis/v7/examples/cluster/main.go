package main

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/goignite/v2/config"
	h "github.com/b2wdigital/goignite/v2/health"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
	"github.com/b2wdigital/goignite/v2/redis/v7"
	"github.com/b2wdigital/goignite/v2/redis/v7/ext/health"
)

func main() {

	config.Load()

	logrus.NewLogger()

	var err error

	healthIntegrator := health.NewDefaultClusterIntegrator()

	_, err = redis.NewDefaultClusterClient(context.Background(), healthIntegrator.Register)
	if err != nil {
		log.Error(err)
	}

	all := h.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
