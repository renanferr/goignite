package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
	gimongo "github.com/b2wdigital/goignite/v2/mongo/v1"
	"github.com/b2wdigital/goignite/v2/mongo/v1/ext/newrelic/v3"
)

func main() {

	giconfig.Load()
	gilogrus.NewLogger()

	gimongo.NewDefaultConn(context.Background(), newrelic.Register)
}
