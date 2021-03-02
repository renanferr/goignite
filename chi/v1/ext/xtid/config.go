package xtid

import (
	gichi "github.com/b2wdigital/goignite/chi/v1"
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	enabled = gichi.ExtRoot + ".xtid.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable xtid middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
