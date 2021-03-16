package gizap

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	ConsoleFormatter = "gi.log.zap.console.formatter"
	FileFormatter    = "gi.log.zap.file.formatter"
)

func init() {

	giconfig.Add(ConsoleFormatter, "TEXT", "formatter TEXT/JSON")
	giconfig.Add(FileFormatter, "TEXT", "formatter TEXT/JSON")
}
