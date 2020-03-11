package zap

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	ConsoleFormatter = "log.zap.console.formatter"
	FileFormatter    = "log.file.console.formatter"
)

func init() {

	log.Println("getting configurations for zap")

	config.Add(ConsoleFormatter, "TEXT", "formatter TEXT/JSON")
	config.Add(FileFormatter, "TEXT", "formatter TEXT/JSON")
}
