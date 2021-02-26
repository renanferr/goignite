package gipromecho

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giprometheus "github.com/b2wdigital/goignite/prometheus/v1"
)

const (
	ConfigRoot = giprometheus.ConfigRoot + ".ext.echo"
	enabled    = ConfigRoot + ".enabled"
	route      = ConfigRoot + ".route"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable echo integration")
	giconfig.Add(route, "/metrics", "define prometheus metrics url")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}

func getRoute() string {
	return giconfig.String(route)
}
