package main

import (
	"context"

	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/elasticsearch/v8"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/logrus/v1"
)

func main() {

	config.Load()

	logrus.NewLogger()

	client, err := elasticsearch.NewDefaultClient(context.Background())
	if err != nil {
		log.Panic(err)
	}

	ping, err := client.Ping(client.Ping.WithPretty())
	if err != nil {
		log.Panic(err)
	}

	log.Infof("status: %v", ping.StatusCode)

}
