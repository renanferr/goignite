package config

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Level          = "logging.level"
	ConsoleEnabled = "logging.console.enabled"
	FileEnabled    = "logging.file.enabled"
	FilePath       = "logging.file.path"
	FileName       = "logging.file.name"
)

func init() {

	log.Println("getting configurations for logging")

	config.Add(Level, "INFO", "log level")
	config.Add(ConsoleEnabled, true, "enable/disable console logging")
	config.Add(FileEnabled, false, "enable/disable file logging")
	config.Add(FilePath, "/tmp", "log path")
	config.Add(FileName, "application.log", "log filename")

}
