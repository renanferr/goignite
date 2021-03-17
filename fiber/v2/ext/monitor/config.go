package monitor

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/fiber/v2"
)

const (
	enabled = fiber.ExtRoot + ".monitor.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable monitor middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
