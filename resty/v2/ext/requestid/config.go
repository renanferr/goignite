package requestid

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/resty/v2"
)

const (
	root    = resty.ExtRoot + ".requestid"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable requestId")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
