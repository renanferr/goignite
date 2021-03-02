package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/config"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	giredis "github.com/b2wdigital/goignite/redis/v7"
	"github.com/b2wdigital/goignite/redis/v7/ext/health"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	var err error

	healthIntegrator := health.NewDefaultClientIntegrator()

	_, err = giredis.NewDefaultClient(context.Background(), healthIntegrator.Register)
	if err != nil {
		gilog.Error(err)
	}

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
