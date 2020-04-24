package fasthttp

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	Name                          = "transport.client.fasthttp.name"
	NoDefaultUserAgentHeader      = "transport.client.fasthttp.nodefaultuseragentheader"
	MaxConnsPerHost               = "transport.client.fasthttp.maxconnsperhost"
	ReadBufferSize                = "transport.client.fasthttp.readbuffersize"
	WriteBufferSize               = "transport.client.fasthttp.writebuffersize"
	ReadTimeout                   = "transport.client.fasthttp.readtimeout"
	WriteTimeout                  = "transport.client.fasthttp.writetimeout"
	MaxIdleConnDuration           = "transport.client.fasthttp.maxidleconnduration"
	DisableHeaderNamesNormalizing = "transport.client.fasthttp.disableheadernamesnormalizing"
	DialDualStack                 = "transport.client.fasthttp.dialdualstack"
	MaxResponseBodySize           = "transport.client.fasthttp.maxresponsebodysize"
	MaxIdemponentCallAttempts     = "transport.client.fasthttp.maxidemponentcallattempts"
)

func init() {

	log.Println("getting configurations for fasthttp client")

	config.Add(Name, "", "used in User-Agent request header")
	config.Add(NoDefaultUserAgentHeader, false, "User-Agent header to be excluded from the Request")
	config.Add(MaxConnsPerHost, 512, "the maximum number of concurrent connections")
	config.Add(ReadBufferSize, 0, "per-connection buffer size for responses' reading")
	config.Add(WriteBufferSize, 0, "per-connection buffer size for requests' writing")
	config.Add(ReadTimeout, 1000*time.Hour, "maximum duration for full response reading (including body).")
	config.Add(WriteTimeout, 1000*time.Hour, "maximum duration for full request writing (including body)")
	config.Add(MaxIdleConnDuration, 10*time.Second, "the default duration before idle keep-alive")
	config.Add(DisableHeaderNamesNormalizing, false, "header names are passed as-is without normalization")
	config.Add(DialDualStack, "", "attempt to connect to both ipv4 and ipv6 addresses if set to true")
	config.Add(MaxResponseBodySize, 52428800, "maximum response body size")
	config.Add(MaxIdemponentCallAttempts, 5, "maximum number of attempts for idempotent calls")
}
