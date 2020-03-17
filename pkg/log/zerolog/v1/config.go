package zerolog

import (
	"github.com/b2wdigital/goignite/pkg/config"
	"log"
)

const (
	Formatter = "log.zerolog.formatter"
)

func init() {

	log.Println("getting configurations for zerolog")

	config.Add(Formatter, "TEXT", "formatter TEXT/JSON")
}
