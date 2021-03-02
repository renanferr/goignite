package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/config"
	gielasticsearch "github.com/b2wdigital/goignite/elasticsearch/v8"
	gielasticsearchhealth "github.com/b2wdigital/goignite/elasticsearch/v8/ext/health"
	gihealth "github.com/b2wdigital/goignite/health"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	o, _ := gielasticsearchhealth.DefaultOptions()
	integrator := gielasticsearchhealth.NewIntegrator(o)

	_, err := gielasticsearch.NewDefaultClient(context.Background(), integrator.Register)
	if err != nil {
		gilog.Error(err)
	}

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
