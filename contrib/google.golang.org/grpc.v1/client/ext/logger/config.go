package logger

import (
	"github.com/b2wdigital/goignite/v2/contrib/google.golang.org/grpc.v1/client"
	"github.com/b2wdigital/goignite/v2/core/config"
)

const (
	root    = client.ExtRoot + ".logger"
	enabled = root + ".enabled"
)

func init() {
	config.Add(enabled, true, "enable/disable logger")
}

func IsEnabled() bool {
	return config.Bool(enabled)
}
