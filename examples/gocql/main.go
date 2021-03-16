package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gigocql "github.com/b2wdigital/goignite/v2/gocql/v0"
	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
	"github.com/b2wdigital/goignite/v2/nats/v1/ext/health"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	options, _ := health.DefaultOptions()
	integrator := health.NewIntegrator(options)

	session, err := gigocql.NewDefaultSession(context.Background(), integrator.Register)
	defer session.Close()

	if err != nil {
		panic(err)
	}

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
