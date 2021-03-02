package newrelic

import (
	gichi "github.com/b2wdigital/goignite/chi/v1"
	giconfig "github.com/b2wdigital/goignite/config"
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

func isEnabled() bool {
	return giconfig.Bool(enabled)
}

func isWebResponseEnabled() bool {
	return giconfig.Bool(enabled)
}
