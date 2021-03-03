package gilog

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root           = "gi.log"
	ConsoleEnabled = root + ".console.enabled"
	ConsoleLevel   = root + ".console.level"
	fileRoot       = root + ".file"
	FileEnabled    = fileRoot + ".enabled"
	FileLevel      = fileRoot + ".level"
	FilePath       = fileRoot + ".path"
	FileName       = fileRoot + ".name"
	FileMaxSize    = fileRoot + ".maxsize"
	FileCompress   = fileRoot + ".compress"
	FileMaxAge     = fileRoot + ".maxage"
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
