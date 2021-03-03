package prometheus

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	ConfigRoot = giecho.ExtRoot + ".prometheus"
	enabled    = ConfigRoot + ".enabled"
	route      = ConfigRoot + ".route"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable prometheus integration")
	giconfig.Add(route, "/metrics", "define prometheus metrics url")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetRoute() string {
	return giconfig.String(route)
}
