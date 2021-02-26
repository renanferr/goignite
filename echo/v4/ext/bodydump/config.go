package bodydump

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	enabled = giecho.MiddlewareRoot + ".bodydump.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable body dump middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
