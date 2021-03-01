package bodylimit

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	enabled = giecho.ExtRoot + ".bodylimit.enabled"
	size    = giecho.ExtRoot + ".bodylimit.size"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable body limit middleware")
	giconfig.Add(size, "8M", "body limit size")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}

func getSize() string {
	return giconfig.String(size)
}
