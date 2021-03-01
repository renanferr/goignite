package newrelic

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gifiber "github.com/b2wdigital/goignite/fiber/v2"
)

const (
	enabled = gifiber.ExtRoot + "newrelic.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable fiber integration")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
