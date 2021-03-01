package health

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gimongo "github.com/b2wdigital/goignite/mongo/v1"
)

const (
	root        = gimongo.ExtRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	giconfig.Add(name, "mongo", "health name")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
