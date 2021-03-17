package swagger

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/echo/v4"
)

const (
	root    = echo.ExtRoot + ".swagger"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable swagger integration")
	config.Add(route, "/swagger/*", "define swagger metrics url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetRoute() string {
	return config.String(route)
}
