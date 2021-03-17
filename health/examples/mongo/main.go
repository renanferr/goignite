package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/config"
	h "github.com/b2wdigital/goignite/v2/health"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
	"github.com/b2wdigital/goignite/v2/mongo/v1"
	"github.com/b2wdigital/goignite/v2/mongo/v1/ext/health"
)

func main() {

	config.Load()
	logrus.NewLogger()

	integrator := health.NewDefaultIntegrator()

	mongo.NewDefaultConn(context.Background(), integrator.Register)

	all := h.CheckAll(context.Background())

	log.Info(all)
}
