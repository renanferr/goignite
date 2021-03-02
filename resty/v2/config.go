package giresty

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root                           = "gi.resty"
	debug                          = root + ".debug"
	closeConnection                = root + ".closeConnection"
	connectionTimeout              = root + ".connectionTimeout"
	keepAlive                      = root + ".keepAlive"
	fallbackDelay                  = root + ".fallbackDelay"
	requestTimeout                 = root + ".requestTimeout"
	transportDisableCompression    = root + ".transport.disableCompression"
	transportDisableKeepAlives     = root + ".transport.disableKeepAlives"
	transportMaxIdleConnsPerHost   = root + ".transport.maxIdleConnsPerHost"
	transportResponseHeaderTimeout = root + ".transport.responseHeaderTimeout"
	transportForceAttemptHTTP2     = root + ".transport.forceAttemptHTTP2"
	transportMaxIdleConns          = root + ".transport.maxIdleConns"
	transportMaxConnsPerHost       = root + ".transport.maxConnsPerHost"
	transportIdleConnTimeout       = root + ".transport.idleConnTimeout"
	transportTLSHandshakeTimeout   = root + ".transport.TLSHandshakeTimeout"
	transportExpectContinueTimeout = root + ".transport.expectContinueTimeout"
	ExtRoot                        = root + ".ext"
)

func init() {

	giconfig.Add(debug, false, "defines global debug request")
	giconfig.Add(closeConnection, false, "defines global http close connection")
	giconfig.Add(connectionTimeout, 3*time.Minute, "defines global http connection timeout")
	giconfig.Add(keepAlive, 30*time.Second, "defines global http keepalive")
	giconfig.Add(fallbackDelay, 300*time.Millisecond, "defines global fallbackDelay")
	giconfig.Add(requestTimeout, 30*time.Second, "defines global http request timeout")
	giconfig.Add(transportDisableCompression, false, "enabled/disable transport compression")
	giconfig.Add(transportDisableKeepAlives, false, "enabled/disable transport keep alives")
	giconfig.Add(transportMaxIdleConnsPerHost, 2, "define transport max idle conns per host")
	giconfig.Add(transportResponseHeaderTimeout, 2*time.Second, "define transport response header timeout")
	giconfig.Add(transportForceAttemptHTTP2, true, "define transport force attempt http2")
	giconfig.Add(transportMaxIdleConns, 100, "define transport max idle conns")
	giconfig.Add(transportMaxConnsPerHost, 100, "define transport max conns per host")
	giconfig.Add(transportIdleConnTimeout, 90*time.Second, "define transport idle conn timeout")
	giconfig.Add(transportTLSHandshakeTimeout, 10*time.Second, "define transport TLS handshake timeout")
	giconfig.Add(transportExpectContinueTimeout, 1*time.Second, "define transport expect continue timeout")
}
