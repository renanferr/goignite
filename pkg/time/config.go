package time

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const FormatTimestamp = "time.timestamp"

func init() {

	log.Println("getting configurations for time")

	config.Add(FormatTimestamp, "2006/01/02 15:04:05.000", "timestamp format")
}
