package ginewrelic

import (
	"log"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ConfigRoot                      = "gi.newrelic"
	AppName                         = ConfigRoot + ".appName"
	License                         = ConfigRoot + ".license"
	Enabled                         = ConfigRoot + ".enabled"
	TracerEnabled                   = ConfigRoot + ".tracerEnabled"
	Labels                          = ConfigRoot + ".labels"
	ServerlessModeEnabled           = ConfigRoot + ".serverless.enabled"
	ServerlessModeAccountID         = ConfigRoot + ".serverless.accountid"
	ServerlessModeTrustedAccountKey = ConfigRoot + ".serverless.trustedAccountKey"
	ServerlessModePrimaryAppID      = ConfigRoot + ".serverless.primaryAppId"
	ServerlessModeApdexThreshold    = ConfigRoot + ".serverless.apdexThreshold"
	ErrorCollectorIgnoreStatusCodes = ConfigRoot + ".serverless.errorCollector.ignoreStatusCodes"
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
	giconfig.Add(ErrorCollectorIgnoreStatusCodes,
		[]int{
			http.StatusBadRequest,
			http.StatusUnauthorized,
			http.StatusPaymentRequired,
			http.StatusForbidden,
			http.StatusNotFound,
			http.StatusMethodNotAllowed,
			http.StatusNotAcceptable,
			http.StatusProxyAuthRequired,
			http.StatusRequestTimeout,
			http.StatusConflict,
			http.StatusGone,
			http.StatusLengthRequired,
			http.StatusPreconditionFailed,
			http.StatusRequestEntityTooLarge,
			http.StatusRequestURITooLong,
			http.StatusUnsupportedMediaType,
			http.StatusRequestedRangeNotSatisfiable,
			http.StatusExpectationFailed,
			http.StatusTeapot,
			http.StatusMisdirectedRequest,
			http.StatusUnprocessableEntity,
			http.StatusLocked,
			http.StatusFailedDependency,
			http.StatusTooEarly,
			http.StatusUpgradeRequired,
			http.StatusPreconditionRequired,
			http.StatusTooManyRequests,
			http.StatusRequestHeaderFieldsTooLarge,
			http.StatusUnavailableForLegalReasons,
		},
	"newrelic serverless mode apdex threshold")
}
