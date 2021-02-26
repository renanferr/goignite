package gipprofecho

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gipprof "github.com/b2wdigital/goignite/pprof"
)

const (
	ConfigRoot = gipprof.ConfigRoot + ".ext.echo"
	enabled    = ConfigRoot + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable echo integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
