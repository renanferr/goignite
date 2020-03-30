package newrelic

import "github.com/b2wdigital/goignite/pkg/config"

const (
	NewRelicAppName                  = "transport.client.newrelic.appname"
	NewRelicLicense                  = "transport.client.newrelic.license"
	NewRelicEnabled                  = "transport.client.newrelic.enabled"
	NewRelicDistributedTracerEnabled = "transport.client.newrelic.tracerenabled"
)

func init() {
	config.Add(NewRelicAppName, "", "application name for new relic")
	config.Add(NewRelicLicense, "", "new relic license key")
	config.Add(NewRelicEnabled, false, "enables new relic")
	config.Add(NewRelicDistributedTracerEnabled, false, "enables new relic distributed tracer")
}
