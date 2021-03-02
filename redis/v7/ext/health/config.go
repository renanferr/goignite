package health

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giredis "github.com/b2wdigital/goignite/redis/v7"
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
