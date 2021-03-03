package newrelic

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	ginats "github.com/b2wdigital/goignite/v2/nats/v1"
)

const (
	root    = ginats.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
