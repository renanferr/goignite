package girestylogger

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giresty "github.com/b2wdigital/goignite/v2/resty/v2"
)

const (
	root    = giresty.ExtRoot + ".logger"
	enabled = root + ".enabled"
	level   = root + ".level"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable logger")
	giconfig.Add(level, "INFO", "sets log level INFO/DEBUG/TRACE")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func Level() string {
	return giconfig.String(level)
}
