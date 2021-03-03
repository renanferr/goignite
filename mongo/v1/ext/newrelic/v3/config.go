package newrelic

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gimongo "github.com/b2wdigital/goignite/v2/mongo/v1"
)

const (
	root    = gimongo.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable mongo integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
