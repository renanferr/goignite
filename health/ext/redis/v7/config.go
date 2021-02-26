package gihealthredis

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ConfigRoot  = "gi.health.ext.redis"
	name        = ConfigRoot + ".name"
	description = ConfigRoot + ".description"
	required    = ConfigRoot + ".required"
	enabled     = ConfigRoot + ".enabled"
)

func init() {

	giconfig.Add(name, "redis", "health name")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
