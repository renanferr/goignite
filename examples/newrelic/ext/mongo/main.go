package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
	"github.com/b2wdigital/goignite/mongo/v1/ext/newrelic/v3"
)

func main() {

	giconfig.Load()
	gilogrus.NewLogger()

	gimongo.NewDefaultClient(context.Background(), newrelic.Register)
}
