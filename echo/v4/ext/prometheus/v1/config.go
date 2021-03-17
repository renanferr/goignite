package prometheus

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/echo/v4"
)

const (
	ConfigRoot = echo.ExtRoot + ".prometheus"
	enabled    = ConfigRoot + ".enabled"
	route      = ConfigRoot + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable prometheus integration")
	config.Add(route, "/metrics", "define prometheus metrics url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func GetRoute() string {
	return config.String(route)
}
