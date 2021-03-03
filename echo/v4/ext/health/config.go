package health

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giecho "github.com/b2wdigital/goignite/v2/echo/v4"
)

const (
	root    = giecho.ExtRoot + ".health"
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
