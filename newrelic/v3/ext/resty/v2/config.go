package giresty

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
	ginewrelic "github.com/b2wdigital/goignite/newrelic/v3"
)

const (
	ConfigRoot = ginewrelic.ConfigRoot + ".ext.resty"
	enabled = ConfigRoot + ".enabled"
)

func init() {

	log.Println("getting configurations for newrelic resty integration")

	giconfig.Add(enabled, true, "enable/disable resty integration")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
