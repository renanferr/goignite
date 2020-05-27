package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/config"
	gihealth "github.com/b2wdigital/goignite/health"
	gihealthredis "github.com/b2wdigital/goignite/health/ext/redis/v7"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	giredis "github.com/b2wdigital/goignite/redis/v7"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	var err error
	var o *gihealthredis.Options

	o, err = gihealthredis.DefaultOptions()
	if err != nil {
		gilog.Error(err)
	}

	err = gihealthredis.ClusterIntegrate(o)
	if err != nil {
		gilog.Error(err)
	}

	_, err = giredis.NewDefaultClusterClient(context.Background())
	if err != nil {
		gilog.Error(err)
	}

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
