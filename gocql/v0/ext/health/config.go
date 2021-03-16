package gigocqlhealth

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gigocql "github.com/b2wdigital/goignite/v2/gocql/v0"
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
