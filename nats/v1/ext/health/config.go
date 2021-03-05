package ginatshealth

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	ginats "github.com/b2wdigital/goignite/v2/nats/v1"
)

const (
	root        = ginats.ExtRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {
	giconfig.Add(name, "nats", "health name")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
