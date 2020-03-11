package nats

import (
	"log"
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/nats-io/nats.go"
)

const (
	MaxReconnects     = "pubsub.client.nats.maxreconnects"
	ReconnectWait     = "pubsub.client.nats.reconnectwait"
	HealthEnabled     = "pubsub.client.nats.health.enabled"
	HealthDescription = "pubsub.client.nats.health.description"
	HealthRequired    = "pubsub.client.nats.health.required"
	Url = "pubsub.client.nats.url"
)

func init() {

	log.Println("getting default configurations for nats")

	config.Add(MaxReconnects, 1000, "define max reconnects to nats server")
	config.Add(ReconnectWait, 1*time.Second, "define reconnects waiting before reconnect to nats server")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, "default connection", "define health description")
	config.Add(Url, nats.DefaultURL, "define nats server url")

}
