package time

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const FormatTimestamp = "gi.time.timestamp"

func init() {

	log.Println("getting configurations for time")

	giconfig.Add(FormatTimestamp, "2006/01/02 15:04:05.000", "timestamp format")
}
