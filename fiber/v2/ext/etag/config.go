package etag

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	enabled = fiber.ExtRoot + ".etag.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable etag middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
