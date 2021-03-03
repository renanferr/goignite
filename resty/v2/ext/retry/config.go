package retry

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	giresty "github.com/b2wdigital/goignite/resty/v2"
)

const (
	root        = giresty.ExtRoot + ".retry"
	enabled     = root + ".enabled"
	count       = root + ".count"
	waitTime    = root + ".waitTime"
	maxWaitTime = root + ".maxWaitTime"
)

func init() {
	giconfig.Add(enabled, true, "enable/disable retry")
	giconfig.Add(count, 0, "defines global max http retries")
	giconfig.Add(waitTime, 200*time.Millisecond, "defines global retry wait time")
	giconfig.Add(maxWaitTime, 2*time.Second, "defines global max retry wait time")
}

func isEnabled() bool {
	return giconfig.Bool(enabled)
}
