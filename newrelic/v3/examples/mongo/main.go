package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
	gimongo "github.com/b2wdigital/goignite/v2/mongo/v1"
	newrelic "github.com/b2wdigital/goignite/v2/mongo/v1/ext/newrelic/v3"
)

func main() {

	giconfig.Load()
	gilogrus.NewLogger()

	gimongo.NewDefaultConn(context.Background(), newrelic.Register)
}
