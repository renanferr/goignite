package realip

import (
	"github.com/b2wdigital/goignite/v2/chi/v5"
	"github.com/b2wdigital/goignite/v2/config"
)

const (
	enabled = chi.ExtRoot + ".realip.enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable realip middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
