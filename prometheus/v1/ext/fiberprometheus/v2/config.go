package giprometheusfiber

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
	giprometheus "github.com/b2wdigital/goignite/prometheus/v1"
)

const (
	ConfigRoot = giprometheus.ConfigRoot + ".ext.fiber"
	enabled    = ConfigRoot + ".enabled"
	route      = ConfigRoot + ".route"
)

func init() {

	log.Println("getting configurations for prometheus fiber integration")

	giconfig.Add(enabled, true, "enable/disable fiber integration")
	giconfig.Add(route, "/metrics", "define prometheus metrics url")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetRoute() string {
	return giconfig.String(route)
}
