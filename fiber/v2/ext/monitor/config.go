package monitor

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	enabled = gifiber.ExtRoot + ".monitor.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable monitor middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
