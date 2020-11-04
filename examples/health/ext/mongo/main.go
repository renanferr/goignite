package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gihealth "github.com/b2wdigital/goignite/health"
	gihealthmongo "github.com/b2wdigital/goignite/health/ext/mongo/v1"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
)

func main() {

	giconfig.Load()
	gilogrus.NewLogger()

	options, _ := gihealthmongo.DefaultOptions()
	gihealthmongo.Integrate(options)

	gimongo.NewDefaultClient(context.Background())

	all := gihealth.CheckAll(context.Background())

	gilog.Info(all)
}
