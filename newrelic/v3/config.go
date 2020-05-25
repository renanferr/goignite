package ginewrelic

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ConfigRoot                      = "gi.newrelic"
	AppName                         = ConfigRoot + ".appname"
	License                         = ConfigRoot + ".license"
	Enabled                         = ConfigRoot + ".enabled"
	TracerEnabled                   = ConfigRoot + ".tracerenabled"
	Labels                          = ConfigRoot + ".labels"
	ServerlessModeEnabled           = ConfigRoot + ".serverless.enabled"
	ServerlessModeAccountID         = ConfigRoot + ".serverless.accountid"
	ServerlessModeTrustedAccountKey = ConfigRoot + ".serverless.trustedaccountkey"
	ServerlessModePrimaryAppID      = ConfigRoot + ".serverless.primaryappid "
	ServerlessModeApdexThreshold    = ConfigRoot + ".serverless.apdexthreshold"
)

func init() {

	log.Println("getting configurations for newrelic")

	giconfig.Add(AppName, "", "application name for newrelic")
	giconfig.Add(License, "", "newrelic license key")
	giconfig.Add(Enabled, false, "enables newrelic")
	giconfig.Add(TracerEnabled, false, "enables newrelic distributed tracer")
	giconfig.Add(Labels, map[string]string{}, "newrelic labels")
	giconfig.Add(ServerlessModeEnabled, false, "enables newrelic serverless mode")
	giconfig.Add(ServerlessModeAccountID, "", "newrelic serverless mode account id")
	giconfig.Add(ServerlessModeTrustedAccountKey, "", "newrelic serverless mode trusted account key")
	giconfig.Add(ServerlessModePrimaryAppID, "", "newrelic serverless mode primary app id")
	giconfig.Add(ServerlessModeApdexThreshold, "", "newrelic serverless mode apdex threshold")
}
