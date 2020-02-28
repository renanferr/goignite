package config

import (
	"log"
	"time"

	"github.com/jpfaria/goignite/pkg/config"
	"github.com/nats-io/nats.go"
)

const (
	Url           = "pubsub.nats.url"
	MaxReconnects = "pubsub.nats.maxreconnects"
	ReconnectWait = "pubsub.nats.reconnectwait"
)

func init() {

	log.Println("getting default configurations for nats")

	config.Add(Url, nats.DefaultURL, "define nats server url")
	config.Add(MaxReconnects, 1000, "define max reconnects to nats server")
	config.Add(ReconnectWait, 1 * time.Second, "define reconnects waiting before reconnect to nats server")
}
