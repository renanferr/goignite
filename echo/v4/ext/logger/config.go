package logger

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	enabled = giecho.ExtRoot + ".logger.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable logging request middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
