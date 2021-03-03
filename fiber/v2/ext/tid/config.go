package tid

import (
	giconfig "github.com/b2wdigital/goignite/config"
	gifiber "github.com/b2wdigital/goignite/fiber/v2"
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
