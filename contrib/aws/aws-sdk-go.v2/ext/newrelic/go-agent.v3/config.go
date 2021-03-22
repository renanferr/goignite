package newrelic

import (
	"github.com/b2wdigital/goignite/v2/contrib/aws/aws-sdk-go.v2"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = aws.ExtRoot + ".newrelic"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable newrelic integration")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
