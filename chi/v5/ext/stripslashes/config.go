package stripslashes

import (
	"github.com/b2wdigital/goignite/v2/chi/v5"
	"github.com/b2wdigital/goignite/v2/config"
)

const (
	root    = chi.ExtRoot + ".stripslashes"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable stripSlashes middleware")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
