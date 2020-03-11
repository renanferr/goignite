package log

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	ConsoleEnabled = "l.console.enabled"
	ConsoleLevel   = "l.console.level"
	FileEnabled    = "l.file.enabled"
	FileLevel      = "l.file.level"
	FilePath       = "l.file.path"
	FileName       = "l.file.name"
	FileMaxSize    = "l.file.maxsize"
	FileCompress   = "l.file.compress"
	FileMaxAge     = "l.file.maxage"
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
