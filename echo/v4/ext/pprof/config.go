package pprof

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	enabled = giecho.IntegrateRoot + ".pprof.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable pprof integration")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
