package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1"
	"github.com/b2wdigital/goignite/v2/contrib/go.mongodb.org/mongo-driver.v1/ext/health"
	"github.com/b2wdigital/goignite/v2/contrib/sirupsen/logrus.v1"
	"github.com/b2wdigital/goignite/v2/core/config"
	h "github.com/b2wdigital/goignite/v2/core/health"
	"github.com/b2wdigital/goignite/v2/core/log"
)

func main() {

	config.Load()
	logrus.NewLogger()

	integrator := health.NewDefaultIntegrator()

	mongo.NewDefaultConn(context.Background(), integrator.Register)

	all := h.CheckAll(context.Background())

	log.Info(all)
}
