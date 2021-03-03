package tid

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
	giresty "github.com/b2wdigital/goignite/v2/resty/v2"
)

const (
	root    = giresty.ExtRoot + ".tid"
	enabled = root + ".enabled"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable tid")
}

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
