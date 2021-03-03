package giresty

import (
	giconfig "github.com/b2wdigital/goignite/config"
	girest "github.com/b2wdigital/goignite/resty/v2"
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
