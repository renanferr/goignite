package newrelic

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giredis "github.com/b2wdigital/goignite/redis/v8"
)

const (
	root    = giredis.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable redis integration")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
