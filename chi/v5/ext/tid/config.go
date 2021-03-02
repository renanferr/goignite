package tid

import (
	gichi "github.com/b2wdigital/goignite/chi/v5"
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	enabled = gichi.ExtRoot + ".tid.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable tid middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
