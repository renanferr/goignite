package zerolog

import "github.com/b2wdigital/goignite/v2/config"

const (
	Formatter = "gi.log.zerolog.formatter"
)

func init() {

	config.Add(Formatter, "TEXT", "formatter TEXT/JSON")
}
