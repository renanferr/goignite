package gifibertid

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	root    = gifiber.ExtRoot + ".tid"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable tid middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
