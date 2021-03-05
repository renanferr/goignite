package gichitid

import (
	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	enabled = gichi.ExtRoot + ".tid.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable tid middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
