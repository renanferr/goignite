package gifasthttp

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	fasthttpClient = "gi.fasthttp.client"

	Name                          = fasthttpClient + ".name"
	NoDefaultUserAgentHeader      = fasthttpClient + ".noDefaultUserAgentHeader"
	MaxConnsPerHost               = fasthttpClient + ".maxConnsPerHost"
	MaxConnWaitTimeout            = fasthttpClient + ".maxConnWaitTimeout"
	ReadBufferSize                = fasthttpClient + ".readBufferSize"
	WriteBufferSize               = fasthttpClient + ".writeBufferSize"
	ReadTimeout                   = fasthttpClient + ".readTimeout"
	WriteTimeout                  = fasthttpClient + ".writeTimeout"
	MaxIdleConnDuration           = fasthttpClient + ".maxIdleConnDuration"
	DisableHeaderNamesNormalizing = fasthttpClient + ".disableHeaderNamesNormalizing"
	DialDualStack                 = fasthttpClient + ".dialDualStack"
	MaxResponseBodySize           = fasthttpClient + ".maxResponseBodySize"
	MaxIdemponentCallAttempts     = fasthttpClient + ".maxIdemponentCallAttempts"
)

func init() {

	giconfig.Add(Name, "", "used in User-Agent request header")
	giconfig.Add(NoDefaultUserAgentHeader, false, "User-Agent header to be excluded from the Request")
	giconfig.Add(MaxConnsPerHost, 512, "the maximum number of concurrent connections")
	giconfig.Add(ReadBufferSize, 0, "per-connection buffer size for responses' reading")
	giconfig.Add(WriteBufferSize, 0, "per-connection buffer size for requests' writing")
	giconfig.Add(MaxConnWaitTimeout, 0*time.Second, "maximum amount of time to wait for a connection to be free")
	giconfig.Add(ReadTimeout, 0*time.Second, "maximum duration for full response reading (including body)")
	giconfig.Add(WriteTimeout, 0*time.Second, "maximum duration for full request writing (including body)")
	giconfig.Add(MaxIdleConnDuration, 10*time.Second, "the default duration before idle keep-alive")
	giconfig.Add(DisableHeaderNamesNormalizing, false, "header names are passed as-is without normalization")
	giconfig.Add(DialDualStack, "", "attempt to connect to both ipv4 and ipv6 addresses if set to true")
	giconfig.Add(MaxResponseBodySize, 52428800, "maximum response body size")
	giconfig.Add(MaxIdemponentCallAttempts, 5, "maximum number of attempts for idempotent calls")
}
