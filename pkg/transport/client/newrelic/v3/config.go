package newrelic

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	AppName       = "transport.client.newrelic.appname"
	License       = "transport.client.newrelic.license"
	Enabled       = "transport.client.newrelic.enabled"
	TracerEnabled = "transport.client.newrelic.tracerenabled"
)

func init() {

	log.Println("getting configurations for newrelic")

	config.Add(AppName, "", "application name for new relic")
	config.Add(License, "", "new relic license key")
	config.Add(Enabled, false, "enables new relic")
	config.Add(TracerEnabled, false, "enables new relic distributed tracer")
}
