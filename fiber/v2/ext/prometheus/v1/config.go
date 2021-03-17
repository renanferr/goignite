package prometheus

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	ConfigRoot = fiber.ExtRoot + ".prometheus"
	enabled    = ConfigRoot + ".enabled"
	route      = ConfigRoot + ".route"
)

func init() {
	config.Add(enabled, true, "enable/disable prometheus integration")
	config.Add(route, "/metrics", "define prometheus metrics url")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}

func getRoute() string {
	return config.String(route)
}
