package status

import (
	gichi "github.com/b2wdigital/goignite/chi/v1"
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root    = gichi.ExtRoot + ".status"
	enabled = root + ".enabled"
	route   = root + ".route"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable status route")
	giconfig.Add(route, "/resource-status", "define status url")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func getRoute() string {
	return giconfig.String(route)
}
