package gihealthmongo

import (
	giconfig "github.com/b2wdigital/goignite/config"

	"log"
)

const (
	ConfigRoot  = "gi.health.ext.mongo"
	name        = ConfigRoot + ".name"
	description = ConfigRoot + ".description"
	required    = ConfigRoot + ".required"
	enabled     = ConfigRoot + ".enabled"
)

func init() {

	log.Println("getting configurations for mongodb health integration")

	giconfig.Add(name, "mongo", "health name")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
