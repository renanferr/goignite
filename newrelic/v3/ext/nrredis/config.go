package ginrredis

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
)

const (
	ConfigRoot = ginewrelic.ConfigRoot + ".ext.redis"
	enabled    = ConfigRoot + ".enabled"
)

func init() {

	log.Println("getting configurations for newrelic redis integration")

	giconfig.Add(enabled, true, "enable/disable redis integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
