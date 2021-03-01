package health

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gigocql "github.com/b2wdigital/goignite/gocql/v0"
)

const (
	root        = gigocql.ExtRoot + ".health"
	name        = root + ".name"
	description = root + ".description"
	required    = root + ".required"
	enabled     = root + ".enabled"
)

func init() {

	giconfig.Add(name, "cassandra", "health name")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
