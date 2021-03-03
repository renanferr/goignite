package prometheus

import (
	gichi "github.com/b2wdigital/goignite/chi/v1"
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root    = gichi.ExtRoot + ".prometheus"
	enabled = root + ".enabled"
	route   = root + ".route"
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
