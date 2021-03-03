package newrelic

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	root                       = giecho.ExtRoot + ".newrelic"
	enabled                    = root + ".enabled"
	middlewareRoot             = root + ".middleware"
	middlewareRequestIDEnabled = middlewareRoot + ".requestId.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic integration")
	giconfig.Add(middlewareRequestIDEnabled, true, "enable/disable request id middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func IsEnabledRequestID() bool {
	return giconfig.Bool(middlewareRequestIDEnabled)
}
