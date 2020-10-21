package gizerolog

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Formatter = "gi.log.zerolog.formatter"
)

func init() {

	log.Println("getting configurations for zerolog")

	giconfig.Add(Formatter, "TEXT", "formatter TEXT/JSON")
}
