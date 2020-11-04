package ginats

import (
	"log"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/nats-io/nats.go"
)

const (
	MaxReconnects   = "gi.nats.maxReconnects"
	ReconnectWait   = "gi.nats.reconnectWait"
	Url             = "gi.nats.url"
	NewRelicEnabled = "gi.nats.newrelic.enabled"
)

func init() {

	log.Println("getting configurations for nats")

	giconfig.Add(MaxReconnects, 1000, "define max reconnects to nats server")
	giconfig.Add(ReconnectWait, 1*time.Second, "define reconnects waiting before reconnect to nats server")
	giconfig.Add(Url, nats.DefaultURL, "define nats server url")
	giconfig.Add(NewRelicEnabled, false, "enable/disable newrelic")

}
