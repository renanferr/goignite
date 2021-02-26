package giprometheusfiber

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giprometheus "github.com/b2wdigital/goignite/prometheus/v1"
)

const (
	ConfigRoot = giprometheus.ConfigRoot + ".ext.fiber"
	enabled    = ConfigRoot + ".enabled"
	route      = ConfigRoot + ".route"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable fiber integration")
	giconfig.Add(route, "/metrics", "define prometheus metrics url")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}

func getRoute() string {
	return giconfig.String(route)
}
