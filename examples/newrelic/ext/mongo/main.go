package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
	ginrmongo "github.com/b2wdigital/goignite/newrelic/v3/ext/nrmongo/v1"
)

func main() {

	giconfig.Load()
	gilogrus.NewLogger()

	ginrmongo.Integrate()

	gimongo.NewDefaultClient(context.Background())
}
