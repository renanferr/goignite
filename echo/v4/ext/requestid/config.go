package requestid

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	enabled = giecho.MiddlewareRoot + ".requestid.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable requestid middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
