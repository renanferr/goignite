package errorhandler

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	enabled = giecho.ExtRoot + ".errorhandler.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable custom error handler")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
