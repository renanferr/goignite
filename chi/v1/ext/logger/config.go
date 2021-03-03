package logger

import (
	gichi "github.com/b2wdigital/goignite/v2/chi/v1"
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	enabled = gichi.ExtRoot + ".logger.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable logger middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
