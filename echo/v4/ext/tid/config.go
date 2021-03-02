package tid

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	enabled = giecho.ExtRoot + ".tid.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable tid middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
