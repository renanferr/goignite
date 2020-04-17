package newrelic

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	AppName                         = "transport.client.newrelic.appname"
	License                         = "transport.client.newrelic.license"
	Enabled                         = "transport.client.newrelic.enabled"
	TracerEnabled                   = "transport.client.newrelic.tracerenabled"
	Labels                          = "transport.client.newrelic.labels"
	ServerlessModeEnabled           = "transport.client.newrelic.serverless.enabled"
	ServerlessModeAccountID         = "transport.client.newrelic.serverless.accountid"
	ServerlessModeTrustedAccountKey = "transport.client.newrelic.serverless.trustedaccountkey"
	ServerlessModePrimaryAppID      = "transport.client.newrelic.serverless.primaryappid "
	ServerlessModeApdexThreshold    = "transport.client.newrelic.serverless.apdexthreshold"
)

func init() {

	log.Println("getting configurations for newrelic")

	config.Add(AppName, "", "application name for newrelic")
	config.Add(License, "", "newrelic license key")
	config.Add(Enabled, false, "enables newrelic")
	config.Add(TracerEnabled, false, "enables newrelic distributed tracer")
	config.Add(Labels, map[string]string{}, "newrelic labels")
	config.Add(ServerlessModeEnabled, false, "enables newrelic serverless mode")
	config.Add(ServerlessModeAccountID, "", "newrelic serverless mode account id")
	config.Add(ServerlessModeTrustedAccountKey, "", "newrelic serverless mode trusted account key")
	config.Add(ServerlessModePrimaryAppID, "", "newrelic serverless mode primary app id")
	config.Add(ServerlessModeApdexThreshold, "", "newrelic serverless mode apdex threshold")
}
