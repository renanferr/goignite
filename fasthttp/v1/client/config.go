package gifasthttp

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"

	"log"
)

const (
	Name                          = "gi.fasthttp.client.name"
	NoDefaultUserAgentHeader      = "gi.fasthttp.client.noDefaultUserAgentHeader"
	MaxConnsPerHost               = "gi.fasthttp.client.maxConnsPerHost"
	ReadBufferSize                = "gi.fasthttp.client.readBufferSize"
	WriteBufferSize               = "gi.fasthttp.client.writeBufferSize"
	ReadTimeout                   = "gi.fasthttp.client.readTimeout"
	WriteTimeout                  = "gi.fasthttp.client.writeTimeout"
	MaxIdleConnDuration           = "gi.fasthttp.client.maxIdleConnDuration"
	DisableHeaderNamesNormalizing = "gi.fasthttp.client.disableHeaderNamesNormalizing"
	DialDualStack                 = "gi.fasthttp.client.dialDualStack"
	MaxResponseBodySize           = "gi.fasthttp.client.maxResponseBodySize"
	MaxIdemponentCallAttempts     = "gi.fasthttp.client.maxIdemponentCallAttempts"
)

func init() {

	log.Println("getting configurations for fasthttp client")

	giconfig.Add(Name, "", "used in User-Agent request header")
	giconfig.Add(NoDefaultUserAgentHeader, false, "User-Agent header to be excluded from the Request")
	giconfig.Add(MaxConnsPerHost, 512, "the maximum number of concurrent connections")
	giconfig.Add(ReadBufferSize, 0, "per-connection buffer size for responses' reading")
	giconfig.Add(WriteBufferSize, 0, "per-connection buffer size for requests' writing")
	giconfig.Add(ReadTimeout, 1000*time.Hour, "maximum duration for full response reading (including body).")
	giconfig.Add(WriteTimeout, 1000*time.Hour, "maximum duration for full request writing (including body)")
	giconfig.Add(MaxIdleConnDuration, 10*time.Second, "the default duration before idle keep-alive")
	giconfig.Add(DisableHeaderNamesNormalizing, false, "header names are passed as-is without normalization")
	giconfig.Add(DialDualStack, "", "attempt to connect to both ipv4 and ipv6 addresses if set to true")
	giconfig.Add(MaxResponseBodySize, 52428800, "maximum response body size")
	giconfig.Add(MaxIdemponentCallAttempts, 5, "maximum number of attempts for idempotent calls")
}
