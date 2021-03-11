package gichilogger

import (
	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root    = gichi.ExtRoot + ".logger"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable logger middleware")
	giconfig.Add(level, "INFO", "sets log level INFO/DEBUG/TRACE")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func Level() string {
	return giconfig.String(level)
}
