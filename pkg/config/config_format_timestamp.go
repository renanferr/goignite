package config

import (
	"log"
)

const FormatTimestamp = "format.timestamp"

func init() {

	log.Println("getting configurations for formats")

	Add(FormatTimestamp, "2006/01/02 15:04:05.000", "timestamp format")
}
