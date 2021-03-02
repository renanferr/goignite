package newrelic

import (
	giconfig "github.com/b2wdigital/goignite/config"
	ginats "github.com/b2wdigital/goignite/nats/v1"
)

const (
	root    = ginats.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
