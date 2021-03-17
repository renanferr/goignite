package datadog

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/echo/v4"
)

const (
	enabled = echo.ExtRoot + ".datadog.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
