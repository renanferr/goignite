package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const LogLevel = "log.level"
const LogConsoleEnabled = "log.console.enabled"
const LogFileEnabled = "log.file.enabled"
const LogFilePath = "log.file.path"
const LogFileName = "log.file.name"

func init() {

	log.Println("getting configurations for logging")

	config.Add(LogLevel, "ERROR", "log level")
	config.Add(LogConsoleEnabled, true, "enable/disable console logging")
	config.Add(LogFileEnabled, false, "enable/disable file logging")
	config.Add(LogFilePath, "/tmp", "log path")
	config.Add(LogFileName, "application.log", "log filename")

}
