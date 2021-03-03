package ginats

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	"github.com/nats-io/nats.go"
)

const (
	root            = "gi.nats"
	maxReconnects   = root + ".maxReconnects"
	reconnectWait   = root + ".reconnectWait"
	url             = root + ".url"
	NewRelicEnabled = root + ".newrelic.enabled"
	ExtRoot         = root + ".Ext"
)

func init() {

	giconfig.Add(maxReconnects, 1000, "define max reconnects to nats server")
	giconfig.Add(reconnectWait, 1*time.Second, "define reconnects waiting before reconnect to nats server")
	giconfig.Add(url, nats.DefaultURL, "define nats server url")
	giconfig.Add(NewRelicEnabled, false, "enable/disable newrelic")

}
