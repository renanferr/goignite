package mongodb

import (
	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	Uri               = "transport.client.mongodb.uri"
	HealthEnabled     = "transport.client.mongodb.health.enabled"
	HealthDescription = "transport.client.mongodb.health.description"
	HealthRequired    = "transport.client.mongodb.health.required"
)

func init() {

	log.Println("getting configurations for mongodb")

	config.Add(Uri, "mongodb://localhost:27017/temp", "define mongodb uri")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, true, "define health description")

}
