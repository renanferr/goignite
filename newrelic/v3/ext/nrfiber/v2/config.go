package ginrfiber

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
)

const (
	configRoot = ginewrelic.ConfigRoot + ".ext.fiber"
	enabled    = configRoot + ".enabled"
)

func init() {

	log.Println("getting configurations for newrelic fiber integration")

	giconfig.Add(enabled, true, "enable/disable fiber integration")

}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
