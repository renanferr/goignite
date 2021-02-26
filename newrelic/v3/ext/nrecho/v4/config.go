package ginrecho

import (
	giconfig "github.com/b2wdigital/goignite/config"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
)

const (
	configRoot                 = ginewrelic.ConfigRoot + ".ext.echo"
	enabled                    = configRoot + ".enabled"
	middlewareRoot             = configRoot + ".middleware"
	middlewareRequestIDEnabled = middlewareRoot + ".requestId.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable echo integration")
	giconfig.Add(middlewareRequestIDEnabled, false, "enable/disable request id middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetMiddlewareRequestIDEnabled() bool {
	return giconfig.Bool(middlewareRequestIDEnabled)
}
