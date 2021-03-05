package giredishealth

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giredis "github.com/b2wdigital/goignite/v2/redis/v8"
)

const (
	root        = giredis.ExtRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {

	giconfig.Add(name, "redis", "health name")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
