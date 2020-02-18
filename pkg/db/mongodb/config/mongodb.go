package config

import (
	"github.com/jpfaria/goignite/pkg/config"

	"log"
)

const Uri = "db.mongodb.uri"
const HealthEnabled = "db.mongodb.health.enabled"
const HealthDescription = "db.mongodb.health.description"
const HealthRequired = "db.mongodb.health.required"

func init() {

	log.Println("getting configurations for mongodb")

	config.Add(Uri, "mongodb://localhost:27017/temp", "define mongodb uri")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, "default connection", "define health description")

}

