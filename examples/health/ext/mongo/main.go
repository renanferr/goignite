package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
	gimongo "github.com/b2wdigital/goignite/v2/mongo/v1"
	health "github.com/b2wdigital/goignite/v2/mongo/v1/ext/health"
)

func main() {

	giconfig.Load()
	gilogrus.NewLogger()

	options, _ := health.DefaultOptions()
	integrator := health.NewIntegrator(options)

	gimongo.NewDefaultConn(context.Background(), integrator.Register)

	all := gihealth.CheckAll(context.Background())

	gilog.Info(all)
}
