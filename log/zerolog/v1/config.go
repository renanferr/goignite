package gizerolog

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Formatter = "gi.log.zerolog.formatter"
)

func init() {

	giconfig.Add(Formatter, "TEXT", "formatter TEXT/JSON")
}
