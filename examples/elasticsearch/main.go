package main

import (
	"context"
	"encoding/json"

	giconfig "github.com/b2wdigital/goignite/config"
	gihealth "github.com/b2wdigital/goignite/health"
	gihealthelasticsearch "github.com/b2wdigital/goignite/health/ext/elasticsearch/v8"
	gilog "github.com/b2wdigital/goignite/log"
	gilogrus "github.com/b2wdigital/goignite/log/logrus/v1"
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	o, _ := gihealthelasticsearch.DefaultOptions()
	gihealthelasticsearch.Integrate(o)

	_, err := elasticsearch.NewDefaultClient()
	if err != nil {
		gilog.Error(err)
	}

	all := gihealth.CheckAll(context.Background())

	j, _ := json.Marshal(all)

	gilog.Info(string(j))

}
