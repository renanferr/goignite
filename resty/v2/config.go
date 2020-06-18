package giresty

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"

	"log"
)

const (
	Debug                          = "gi.resty.debug"
	ConnectionTimeout              = "gi.resty.connectionTimeout"
	KeepAlive                      = "gi.resty.keepAlive"
	RequestTimeout                 = "gi.resty.requestTimeout"
	RetryCount                     = "gi.resty.retry.count"
	RetryWaitTime                  = "gi.resty.retry.waitTime"
	RetryMaxWaitTime               = "gi.resty.retry.maxWaitTime"
	TransportDisableCompression    = "gi.resty.transport.disableCompression"
	TransportDisableKeepAlives     = "gi.resty.transport.disableKeepAlives"
	TransportMaxIdleConnsPerHost   = "gi.resty.transport.maxIdleConnsPerHost"
	TransportResponseHeaderTimeout = "gi.resty.transport.responseHeaderTimeout"
	TransportForceAttemptHTTP2     = "gi.resty.transport.forceAttemptHTTP2"
	TransportMaxIdleConns          = "gi.resty.transport.maxIdleConns"
	TransportMaxConnsPerHost       = "gi.resty.transport.maxConnsPerHost"
	TransportIdleConnTimeout       = "gi.resty.transport.idleConnTimeout"
	TransportTLSHandshakeTimeout   = "gi.resty.transport.TLSHandshakeTimeout"
	TransportExpectContinueTimeout = "gi.resty.transport.expectContinueTimeout"
)

func init() {
	log.Println("getting configurations for resty")

	giconfig.Add(Debug, false, "defines global debug request")
	giconfig.Add(ConnectionTimeout, 3*time.Minute, "defines global http connection timeout")
	giconfig.Add(KeepAlive, 30*time.Second, "defines global http keepalive")
	giconfig.Add(RequestTimeout, 30*time.Second, "defines global http request timeout")
	giconfig.Add(RetryCount, 0, "defines global max http retries")
	giconfig.Add(RetryWaitTime, 200*time.Millisecond, "defines global retry wait time")
	giconfig.Add(RetryMaxWaitTime, 2*time.Second, "defines global max retry wait time")

	giconfig.Add(TransportDisableCompression, false, "enabled/disable transport compression")
	giconfig.Add(TransportDisableKeepAlives, false, "enabled/disable transport keep alives")
	giconfig.Add(TransportMaxIdleConnsPerHost, 2, "define transport max idle conns per host")
	giconfig.Add(TransportResponseHeaderTimeout, 2*time.Second, "define transport response header timeout")
	giconfig.Add(TransportForceAttemptHTTP2, true, "define transport force attempt http2")
	giconfig.Add(TransportMaxIdleConns, 100, "define transport max idle conns")
	giconfig.Add(TransportMaxConnsPerHost, 100, "define transport max conns per host")
	giconfig.Add(TransportIdleConnTimeout, 90*time.Second, "define transport idle conn timeout")
	giconfig.Add(TransportTLSHandshakeTimeout, 10*time.Second, "define transport TLS handshake timeout")
	giconfig.Add(TransportExpectContinueTimeout, 1*time.Second, "define transport expect continue timeout")
}
