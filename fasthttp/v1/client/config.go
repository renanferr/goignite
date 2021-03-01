package gifasthttp

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root                          = "gi.fasthttp.client"
	name                          = root + ".name"
	noDefaultUserAgentHeader      = root + ".noDefaultUserAgentHeader"
	maxConnsPerHost               = root + ".maxConnsPerHost"
	maxConnWaitTimeout            = root + ".maxConnWaitTimeout"
	readBufferSize                = root + ".readBufferSize"
	writeBufferSize               = root + ".writeBufferSize"
	readTimeout                   = root + ".readTimeout"
	writeTimeout                  = root + ".writeTimeout"
	maxIdleConnDuration           = root + ".maxIdleConnDuration"
	disableHeaderNamesNormalizing = root + ".disableHeaderNamesNormalizing"
	dialDualStack                 = root + ".dialDualStack"
	maxResponseBodySize           = root + ".maxResponseBodySize"
	maxIdemponentCallAttempts     = root + ".maxIdemponentCallAttempts"
)

func init() {

	giconfig.Add(name, "", "used in User-Agent request header")
	giconfig.Add(noDefaultUserAgentHeader, false, "User-Agent header to be excluded from the Request")
	giconfig.Add(maxConnsPerHost, 512, "the maximum number of concurrent connections")
	giconfig.Add(readBufferSize, 0, "per-connection buffer size for responses' reading")
	giconfig.Add(writeBufferSize, 0, "per-connection buffer size for requests' writing")
	giconfig.Add(maxConnWaitTimeout, 0*time.Second, "maximum amount of time to wait for a connection to be free")
	giconfig.Add(readTimeout, 0*time.Second, "maximum duration for full response reading (including body)")
	giconfig.Add(writeTimeout, 0*time.Second, "maximum duration for full request writing (including body)")
	giconfig.Add(maxIdleConnDuration, 10*time.Second, "the default duration before idle keep-alive")
	giconfig.Add(disableHeaderNamesNormalizing, false, "header names are passed as-is without normalization")
	giconfig.Add(dialDualStack, "", "attempt to connect to both ipv4 and ipv6 addresses if set to true")
	giconfig.Add(maxResponseBodySize, 52428800, "maximum response body size")
	giconfig.Add(maxIdemponentCallAttempts, 5, "maximum number of attempts for idempotent calls")
}
