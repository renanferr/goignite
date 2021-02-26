package giswaggerecho

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giswagger "github.com/b2wdigital/goignite/swagger/v1"
)

const (
	ConfigRoot = giswagger.ConfigRoot + ".ext.echo"
	enabled    = ConfigRoot + ".enabled"
	route      = ConfigRoot + ".route"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable echo integration")
	giconfig.Add(route, "/swagger", "define swagger metrics url")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}

func GetRoute() string {
	return giconfig.String(route)
}
