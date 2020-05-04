package mongodb

import (
	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	ConfigRoot		  = "transport.client.mongodb"
	Uri               = ConfigRoot + ".uri"
	HealthEnabled     = ConfigRoot + ".enabled"
	HealthDescription = ConfigRoot + ".health.description"
	HealthRequired    = ConfigRoot + ".health.required"
	NewRelicEnabled   = ConfigRoot + ".newrelic.enabled"
)

func init() {

	log.Println("getting configurations for mongodb")

	config.Add(Uri, "mongodb://localhost:27017/temp", "define mongodb uri")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, true, "define health description")
	config.Add(NewRelicEnabled, false, "enable/disable newrelic")

}
