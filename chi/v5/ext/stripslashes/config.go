package gichistripslashes

import (
	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root    = gichi.ExtRoot + ".stripslashes"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable stripSlashes middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
