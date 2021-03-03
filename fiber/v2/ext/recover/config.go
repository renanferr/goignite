package recover

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	enabled = gifiber.ExtRoot + ".recover.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable recover middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
