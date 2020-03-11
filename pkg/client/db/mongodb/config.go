package mongodb

import (
	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	Uri               = "db.mongodb.uri"
	HealthEnabled     = "db.mongodb.health.enabled"
	HealthDescription = "db.mongodb.health.description"
	HealthRequired    = "db.mongodb.health.required"
)

func init() {

	log.Println("getting configurations for mongodb")

	config.Add(Uri, "mongodb://localhost:27017/temp", "define mongodb uri")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, "default connection", "define health description")

}
