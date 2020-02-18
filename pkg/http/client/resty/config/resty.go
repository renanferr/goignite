package config

import (
	"time"

	"github.com/jpfaria/goignite/pkg/config"

	"log"
)

const Debug = "resty.debug"
const RequestTimeout = "resty.request.timeout"
const RetryCount = "resty.retry.count"
const RetryWaitTime = "resty.retry.waittime"
const RetryMaxWaitTime = "resty.retry.maxwaittime"
const HealthEnabled = "resty.health.enabled"
const TransportDisableCompression = "resty.transport.disablecompression"
const TransportDisableKeepAlives = "resty.transport.disablekeepalives"
const TransportMaxIdleConnsPerHost = "resty.transport.maxidleconnsperhost"
const TransportResponseHeaderTimeout = "resty.transport.responseheadertimeout"
const TransportForceAttemptHTTP2 = "resty.transport.forceattempthttp2"
const TransportMaxIdleConns = "resty.transport.maxidleconns"
const TransportMaxConnsPerHost = "resty.transport.maxconnsperhost"
const TransportIdleConnTimeout = "resty.transport.idleconntimeout"
const TransportTLSHandshakeTimeout = "resty.transport.tlshandshaketimeout"
const TransportExpectContinueTimeout = "resty.transport.expectcontinuetimeout"

func init() {
	log.Println("getting default configurations for resty")

	config.Add(Debug, false, "defines global debug request")
	config.Add(RequestTimeout, 2 * time.Second, "defines global http request timeout (ms)")
	config.Add(RetryCount, 0, "defines global max http retries")
	config.Add(RetryWaitTime, 200 * time.Millisecond, "defines global retry wait time (ms)")
	config.Add(RetryMaxWaitTime, 2 * time.Second, "defines global max retry wait time (ms)")
	config.Add(HealthEnabled, true, "enabled/disable health check")

	config.Add(TransportDisableCompression, false, "enabled/disable transport compression")
	config.Add(TransportDisableKeepAlives, false, "enabled/disable transport keep alives")
	config.Add(TransportMaxIdleConnsPerHost, 2, "define transport max idle conns per host")
	config.Add(TransportResponseHeaderTimeout, 2 * time.Second, "define transport response header timeout")
	config.Add(TransportForceAttemptHTTP2, true, "define transport force attempt http2")
	config.Add(TransportMaxIdleConns, 100, "define transport max idle conns")
	config.Add(TransportMaxConnsPerHost, 100, "define transport max conns per host")
	config.Add(TransportIdleConnTimeout, 90 * time.Second, "define transport idle conn timeout")
	config.Add(TransportTLSHandshakeTimeout, 10 * time.Second, "define transport TLS handshake timeout")
	config.Add(TransportExpectContinueTimeout, 1 * time.Second, "define transport expect continue timeout")
}
