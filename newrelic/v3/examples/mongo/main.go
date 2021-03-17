package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
	"github.com/b2wdigital/goignite/v2/mongo/v1"
	"github.com/b2wdigital/goignite/v2/mongo/v1/ext/newrelic/v3"
)

func main() {

	config.Load()
	logrus.NewLogger()

	mongo.NewDefaultConn(context.Background(), newrelic.Register)
}
