package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gielasticsearch "github.com/b2wdigital/goignite/v2/elasticsearch/v8"
	gielasticsearchhealth "github.com/b2wdigital/goignite/v2/elasticsearch/v8/ext/health"
	gihealth "github.com/b2wdigital/goignite/v2/health"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	integrator := gielasticsearchhealth.NewDefaultIntegrator()

	_, err := gielasticsearch.NewDefaultClient(context.Background(), integrator.Register)
	if err != nil {
		gilog.Panic(err)
	}

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
