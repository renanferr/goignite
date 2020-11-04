package gihealthresty

import (
	giconfig "github.com/b2wdigital/goignite/config"

	"log"
)

const (
	ConfigRoot  = "gi.health.ext.resty"
	name        = ConfigRoot + ".name"
	host        = ConfigRoot + ".host"
	endpoint    = ConfigRoot + ".endpoint"
	description = ConfigRoot + ".description"
	required    = ConfigRoot + ".required"
	enabled     = ConfigRoot + ".enabled"
)

func init() {

	log.Println("getting configurations for resty health integration")

	giconfig.Add(name, "rest api", "health name")
	giconfig.Add(host, "", "health host")
	giconfig.Add(endpoint, "/resource-status", "health host")
	giconfig.Add(description, "default connection", "define health description")
	giconfig.Add(required, true, "define health description")
	giconfig.Add(enabled, true, "enable/disable health")
}
