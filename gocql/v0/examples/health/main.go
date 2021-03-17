package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gigocql "github.com/b2wdigital/goignite/v2/gocql/v0"
	gigocqlhealth "github.com/b2wdigital/goignite/v2/gocql/v0/ext/health"
	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	integrator := gigocqlhealth.NewDefaultIntegrator()

	session, err := gigocql.NewDefaultSession(context.Background(), integrator.Register)
	if err != nil {
		panic(err)
	}

	defer session.Close()

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
