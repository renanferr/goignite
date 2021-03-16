package gifx

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root     = "gi.fx"
	logLevel = root + ".log.level"
)

func init() {
	giconfig.Add(logLevel, "DEBUG", "define log level")
}

func LogLevel() string {
	return giconfig.String(logLevel)
}
