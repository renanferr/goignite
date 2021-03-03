package health

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	root    = gifiber.ExtRoot + ".health"
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

func getRoute() string {
	return giconfig.String(route)
}
