package pprof

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gifiber "github.com/b2wdigital/goignite/fiber/v2"
)

const (
	enabled = gifiber.ExtRoot + ".pprof.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable pprof middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
