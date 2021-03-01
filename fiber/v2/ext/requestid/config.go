package requestid

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gifiber "github.com/b2wdigital/goignite/fiber/v2"
)

const (
	enabled = gifiber.ExtRoot + ".requestid.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable requestid middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
