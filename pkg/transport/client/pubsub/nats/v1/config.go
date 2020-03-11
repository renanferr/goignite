package nats

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/nats-io/nats.go"
)

const (
	Url = "pubsub.client.nats.url"
)

func init() {

	log.Println("getting default configurations for nats")

	config.Add(Url, nats.DefaultURL, "define nats server url")

}
