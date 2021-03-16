package giantsnewrelic

import (
	giants "github.com/b2wdigital/goignite/v2/ants/v2"
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root    = giants.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
