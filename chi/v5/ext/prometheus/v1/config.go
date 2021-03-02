package prometheus

import (
	gichi "github.com/b2wdigital/goignite/chi/v5"
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

func isEnabled() bool {
	return giconfig.Bool(enabled)
}

func getRoute() string {
	return giconfig.String(route)
}
