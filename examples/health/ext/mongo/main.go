package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/config"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
	"github.com/b2wdigital/goignite/mongo/v1/ext/health"
)

func main() {

	giconfig.Load()
	gilogrus.NewLogger()

	options, _ := health.DefaultOptions()
	integrator := health.NewIntegrator(options)

	gimongo.NewDefaultClient(context.Background(), integrator.Integrate)

	all := gihealth.CheckAll(context.Background())

	gilog.Info(all)
}
