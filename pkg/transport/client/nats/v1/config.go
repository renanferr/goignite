package nats

import (
	"log"
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/nats-io/nats.go"
)

const (
	MaxReconnects     = "transport.client.nats.maxreconnects"
	ReconnectWait     = "transport.client.nats.reconnectwait"
	HealthEnabled     = "transport.client.nats.health.enabled"
	HealthDescription = "transport.client.nats.health.description"
	HealthRequired    = "transport.client.nats.health.required"
	Url               = "transport.client.nats.url"
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
