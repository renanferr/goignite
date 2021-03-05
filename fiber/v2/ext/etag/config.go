package gifiberetag

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	gifiber "github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	enabled = gifiber.ExtRoot + ".etag.enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable etag middleware")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
