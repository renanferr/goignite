package logger

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giresty "github.com/b2wdigital/goignite/resty/v2"
)

const (
	root    = giresty.ExtRoot + ".logger"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable logger")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
