package health

import (
	gichi "github.com/b2wdigital/goignite/v2/chi/v1"
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root    = gichi.ExtRoot + ".health"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable health route")
	giconfig.Add(route, "/health", "define status url")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetRoute() string {
	return giconfig.String(route)
}
