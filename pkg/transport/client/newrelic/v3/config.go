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
	Labels        = "transport.client.newrelic.labels"
)

func init() {

	log.Println("getting configurations for newrelic")

	config.Add(AppName, "", "application name for newrelic")
	config.Add(License, "", "newrelic license key")
	config.Add(Enabled, false, "enables newrelic")
	config.Add(TracerEnabled, false, "enables newrelic distributed tracer")
	config.Add(Labels, map[string]string{}, "newrelic labels")
}
