package recover

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gifiber "github.com/b2wdigital/goignite/fiber/v2"
)

const (
	enabled = gifiber.ExtRoot + ".recover.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable recover middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
