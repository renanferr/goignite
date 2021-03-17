package newrelic

import (
	"github.com/b2wdigital/goignite/v2/config"
	girest "github.com/b2wdigital/goignite/v2/resty/v2"
)

const (
	ConfigRoot = girest.ExtRoot + ".newrelic"
	enabled    = ConfigRoot + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
