package health

import (
	"github.com/b2wdigital/goignite/v2/config"
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
	config.Add(name, "nats", "health name")
	config.Add(description, "default connection", "define health description")
	config.Add(required, true, "define health description")
	config.Add(enabled, true, "enable/disable health")
}
