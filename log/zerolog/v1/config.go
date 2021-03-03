package gizerolog

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	Formatter = "gi.log.zerolog.formatter"
)

func init() {

	giconfig.Add(Formatter, "TEXT", "formatter TEXT/JSON")
}
