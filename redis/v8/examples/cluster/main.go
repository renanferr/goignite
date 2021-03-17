package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
	giredis "github.com/b2wdigital/goignite/v2/redis/v8"
	health "github.com/b2wdigital/goignite/v2/redis/v8/ext/health"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	var err error

	healthIntegrator := health.NewDefaultClusterIntegrator()

	_, err = giredis.NewDefaultClusterClient(context.Background(), healthIntegrator.Register)
	if err != nil {
		gilog.Error(err)
	}

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
