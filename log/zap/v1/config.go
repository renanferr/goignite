package gizap

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ConsoleFormatter = "gi.log.zap.console.formatter"
	FileFormatter    = "gi.log.zap.file.formatter"
)

func init() {

	log.Println("getting configurations for zap")

	giconfig.Add(ConsoleFormatter, "TEXT", "formatter TEXT/JSON")
	giconfig.Add(FileFormatter, "TEXT", "formatter TEXT/JSON")
}
