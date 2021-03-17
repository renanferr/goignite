package errorhandler

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/echo/v4"
)

const (
	enabled = echo.ExtRoot + ".errorhandler.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable custom error handler")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
