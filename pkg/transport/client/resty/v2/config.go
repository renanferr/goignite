package resty

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	Debug                          = "transport.client.resty.debug"
	RequestTimeout                 = "transport.client.resty.request.timeout"
	RetryCount                     = "transport.client.resty.retry.count"
	RetryWaitTime                  = "transport.client.resty.retry.waittime"
	RetryMaxWaitTime               = "transport.client.resty.retry.maxwaittime"
	TransportDisableCompression    = "transport.client.resty.transport.disablecompression"
	TransportDisableKeepAlives     = "transport.client.resty.transport.disablekeepalives"
	TransportMaxIdleConnsPerHost   = "transport.client.resty.transport.maxidleconnsperhost"
	TransportResponseHeaderTimeout = "transport.client.resty.transport.responseheadertimeout"
	TransportForceAttemptHTTP2     = "transport.client.resty.transport.forceattempthttp2"
	TransportMaxIdleConns          = "transport.client.resty.transport.maxidleconns"
	TransportMaxConnsPerHost       = "transport.client.resty.transport.maxconnsperhost"
	TransportIdleConnTimeout       = "transport.client.resty.transport.idleconntimeout"
	TransportTLSHandshakeTimeout   = "transport.client.resty.transport.tlshandshaketimeout"
	TransportExpectContinueTimeout = "transport.client.resty.transport.expectcontinuetimeout"
)

func init() {
	log.Println("getting default configurations for resty")

	config.Add(Debug, false, "defines global debug request")
	config.Add(RequestTimeout, 2*time.Second, "defines global http request timeout (ms)")
	config.Add(RetryCount, 0, "defines global max http retries")
	config.Add(RetryWaitTime, 200*time.Millisecond, "defines global retry wait time (ms)")
	config.Add(RetryMaxWaitTime, 2*time.Second, "defines global max retry wait time (ms)")

	config.Add(TransportDisableCompression, false, "enabled/disable transport compression")
	config.Add(TransportDisableKeepAlives, false, "enabled/disable transport keep alives")
	config.Add(TransportMaxIdleConnsPerHost, 2, "define transport max idle conns per host")
	config.Add(TransportResponseHeaderTimeout, 2*time.Second, "define transport response header timeout")
	config.Add(TransportForceAttemptHTTP2, true, "define transport force attempt http2")
	config.Add(TransportMaxIdleConns, 100, "define transport max idle conns")
	config.Add(TransportMaxConnsPerHost, 100, "define transport max conns per host")
	config.Add(TransportIdleConnTimeout, 90*time.Second, "define transport idle conn timeout")
	config.Add(TransportTLSHandshakeTimeout, 10*time.Second, "define transport TLS handshake timeout")
	config.Add(TransportExpectContinueTimeout, 1*time.Second, "define transport expect continue timeout")
}
