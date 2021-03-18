package time

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/b2wdigital/goignite/v2/config"
)

const (
	root     = "gi.time"
	format   = root + ".timestamp"
	location = root + ".location"
)

func init() {
	config.Add(format, "2006/01/02 15:04:05.000", "time format")
	config.Add(location, time.UTC.String(), "time location")
}

func Format() string {
	return giconfig.String(format)
}

func Location() *time.Location {
	locStr := giconfig.String(location)
	loc, err := time.LoadLocation(locStr)
	if err != nil {
		panic(err)
	}
	return loc
}
