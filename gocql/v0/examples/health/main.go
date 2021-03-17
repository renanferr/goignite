package main

import (
	"context"
	"encoding/json"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/gocql/v0"
	gigocqlhealth "github.com/b2wdigital/goignite/v2/gocql/v0/ext/health"
	"github.com/b2wdigital/goignite/v2/health"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

	integrator := gigocqlhealth.NewDefaultIntegrator()

	session, err := gocql.NewDefaultSession(context.Background(), integrator.Register)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	all := health.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	log.Info(string(j))

}
