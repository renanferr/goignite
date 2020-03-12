package log

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	ConsoleEnabled = "log.console.enabled"
	ConsoleLevel   = "log.console.level"
	FileEnabled    = "log.file.enabled"
	FileLevel      = "log.file.level"
	FilePath       = "log.file.path"
	FileName       = "log.file.name"
	FileMaxSize    = "log.file.maxsize"
	FileCompress   = "log.file.compress"
	FileMaxAge     = "log.file.maxage"
)

func init() {

	log.Println("getting configurations for logging")

	config.Add(ConsoleEnabled, true, "enable/disable console logging")
	config.Add(ConsoleLevel, "INFO", "console l level")
	config.Add(FileEnabled, false, "enable/disable file logging")
	config.Add(FileLevel, "INFO", "console l level")
	config.Add(FilePath, "/tmp", "l path")
	config.Add(FileName, "application.l", "l filename")
	config.Add(FileMaxSize, 100, "l file max size (MB)")
	config.Add(FileCompress, true, "l file compress")
	config.Add(FileMaxAge, 28, "l file max age (days)")

}
