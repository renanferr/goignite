package gilog

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	ConsoleEnabled = "gi.log.console.enabled"
	ConsoleLevel   = "gi.log.console.level"
	FileEnabled    = "gi.log.file.enabled"
	FileLevel      = "gi.log.file.level"
	FilePath       = "gi.log.file.path"
	FileName       = "gi.log.file.name"
	FileMaxSize    = "gi.log.file.maxsize"
	FileCompress   = "gi.log.file.compress"
	FileMaxAge     = "gi.log.file.maxage"
)

func init() {

	giconfig.Add(ConsoleEnabled, true, "enable/disable console logging")
	giconfig.Add(ConsoleLevel, "INFO", "console log level")
	giconfig.Add(FileEnabled, false, "enable/disable file logging")
	giconfig.Add(FileLevel, "INFO", "console log level")
	giconfig.Add(FilePath, "/tmp", "log path")
	giconfig.Add(FileName, "application.l", "log filename")
	giconfig.Add(FileMaxSize, 100, "log file max size (MB)")
	giconfig.Add(FileCompress, true, "log file compress")
	giconfig.Add(FileMaxAge, 28, "log file max age (days)")

}
