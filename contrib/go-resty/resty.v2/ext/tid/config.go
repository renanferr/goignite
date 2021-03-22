package tid

import (
	"github.com/b2wdigital/goignite/v2/contrib/go-resty/resty.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = resty.ExtRoot + ".tid"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable tid")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
