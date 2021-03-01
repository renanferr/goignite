package gzip

import (
	giconfig "github.com/b2wdigital/goignite/config"
	giecho "github.com/b2wdigital/goignite/echo/v4"
)

const (
	enabled = giecho.ExtRoot + ".gzip.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable gzip middleware")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
