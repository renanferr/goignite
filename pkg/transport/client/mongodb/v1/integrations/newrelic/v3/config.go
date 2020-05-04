package newrelic

import (
	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/transport/client/mongodb/v1"

	"log"
)

const (
	Enabled = mongodb.ConfigRoot + ".integration.newrelic.enabled"
)

func init() {

	log.Println("getting configurations for newrelic mongodb integration")

	config.Add(Enabled, false, "enable/disable newrelic")
}
