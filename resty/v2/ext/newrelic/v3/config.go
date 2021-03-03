package newrelic

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	girest "github.com/b2wdigital/goignite/v2/resty/v2"
)

const (
	ConfigRoot = girest.ExtRoot + ".newrelic"
	enabled    = ConfigRoot + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
