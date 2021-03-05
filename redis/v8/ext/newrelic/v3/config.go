package giredisnewrelic

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giredis "github.com/b2wdigital/goignite/v2/redis/v8"
)

const (
	root    = giredis.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable redis integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
