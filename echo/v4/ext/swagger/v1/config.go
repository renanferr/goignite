package giswaggerecho

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	root    = giecho.ExtRoot + ".swagger"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable swagger integration")
	giconfig.Add(route, "/swagger", "define swagger metrics url")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}

func getRoute() string {
	return giconfig.String(route)
}
