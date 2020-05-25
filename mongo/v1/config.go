package gimongo

import (
	giconfig "github.com/b2wdigital/goignite/config"

	"log"
)

const (
	ConfigRoot = "gi.mongo"
	Uri        = ConfigRoot + ".uri"
)

func init() {

	log.Println("getting configurations for mongodb")

	giconfig.Add(Uri, "mongodb://localhost:27017/temp", "define mongodb uri")
}
