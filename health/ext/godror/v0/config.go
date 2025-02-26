package gihealthgodror

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ConfigRoot  = "gi.health.ext.godror"
	name        = ConfigRoot + ".name"
	description = ConfigRoot + ".description"
	required    = ConfigRoot + ".required"
	enabled     = ConfigRoot + ".enabled"
)

func init() {

	giconfig.Add(name, "oracle", "health name")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
