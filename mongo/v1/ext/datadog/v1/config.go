package datadog

import (
	"github.com/b2wdigital/goignite/v2/config"
	"github.com/b2wdigital/goignite/v2/mongo/v1"
)

const (
	root    = mongo.ExtRoot + ".datadog"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable datadog integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
