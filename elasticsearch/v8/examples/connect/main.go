package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gielasticsearch "github.com/b2wdigital/goignite/v2/elasticsearch/v8"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	giconfig.Load()

	gilogrus.NewLogger()

	client, err := gielasticsearch.NewDefaultClient(context.Background())
	if err != nil {
		gilog.Panic(err)
	}

	ping, err := client.Ping(client.Ping.WithPretty())
	if err != nil {
		gilog.Panic(err)
	}

	gilog.Infof("status: %v", ping.StatusCode)

}
