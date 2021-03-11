package giechologger

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
)

const (
	root    = giecho.ExtRoot + ".logger"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable logging request middleware")
	giconfig.Add(level, "INFO", "sets log level INFO/DEBUG/TRACE")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func Level() string {
	return giconfig.String(level)
}
