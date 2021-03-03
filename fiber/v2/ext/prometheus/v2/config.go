package prometheus

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gifiber "github.com/b2wdigital/goignite/fiber/v2"
)

const (
	ConfigRoot = gifiber.ExtRoot + ".prometheus"
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

func getRoute() string {
	return giconfig.String(route)
}
