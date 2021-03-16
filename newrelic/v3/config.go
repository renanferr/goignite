package ginewrelic

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root                            = "gi.newrelic"
	appName                         = root + ".appName"
	license                         = root + ".license"
	enabled                         = root + ".enabled"
	tracerEnabled                   = root + ".tracerEnabled"
	labels                          = root + ".labels"
	serverlessModeEnabled           = root + ".serverless.enabled"
	serverlessModeAccountID         = root + ".serverless.accountid"
	serverlessModeTrustedAccountKey = root + ".serverless.trustedAccountKey"
	serverlessModePrimaryAppID      = root + ".serverless.primaryAppId"
	serverlessModeApdexThreshold    = root + ".serverless.apdexThreshold"
	errorCollectorIgnoreStatusCodes = root + ".errorCollector.ignoreStatusCodes"
)

func init() {

	giconfig.Add(appName, "", "application name for newrelic")
	giconfig.Add(license, "", "newrelic license key")
	giconfig.Add(enabled, false, "enables newrelic")
	giconfig.Add(tracerEnabled, false, "enables newrelic distributed tracer")
	giconfig.Add(labels, map[string]string{}, "newrelic labels")
	giconfig.Add(serverlessModeEnabled, false, "enables newrelic serverless mode")
	giconfig.Add(serverlessModeAccountID, "", "newrelic serverless mode account id")
	giconfig.Add(serverlessModeTrustedAccountKey, "", "newrelic serverless mode trusted account key")
	giconfig.Add(serverlessModePrimaryAppID, "", "newrelic serverless mode primary app id")
	giconfig.Add(serverlessModeApdexThreshold, "", "newrelic serverless mode apdex threshold")
	giconfig.Add(errorCollectorIgnoreStatusCodes,
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
