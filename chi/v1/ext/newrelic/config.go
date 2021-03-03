package newrelic

import (
	gichi "github.com/b2wdigital/goignite/v2/chi/v1"
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root               = gichi.ExtRoot + ".newrelic"
	enabled            = root + ".enabled"
	webResponseEnabled = root + ".webresponse.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic middleware")
	giconfig.Add(webResponseEnabled, true, "enable/disable newrelic web response")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func IsWebResponseEnabled() bool {
	return giconfig.Bool(enabled)
}
