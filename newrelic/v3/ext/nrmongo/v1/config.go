package ginrmongo

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
)

const (
	ConfigRoot = ginewrelic.ConfigRoot + ".ext.mongo"
	enabled    = ConfigRoot + ".enabled"
)

func init() {

	log.Println("getting configurations for newrelic mongodb integration")

	giconfig.Add(enabled, true, "enable/disable mongo integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
