package gihttp

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root                  = "gi.http.client"
	maxIdleConnPerHost    = root + ".maxIdleConnPerHost"
	maxIdleConn           = root + ".maxIdleConn"
	maxConnsPerHost       = root + ".maxConnsPerHost"
	idleConnTimeout       = root + ".idleConnTimeout"
	disableKeepAlives     = root + ".disableKeepAlives"
	forceHTTP2            = root + ".forceHTTP2"
	tlsHandshakeTimeout   = root + ".TLSHandshakeTimeout"
	timeout               = root + ".timeout"
	keepAlive             = root + ".keepAlive"
	expectContinueTimeout = root + ".expectContinueTimeout"
	dualStack             = root + ".dualStack"
)

func init() {

	giconfig.Add(maxIdleConnPerHost, 1, "http max idle connections per host")
	giconfig.Add(maxIdleConn, 100, "http max idle connections")
	giconfig.Add(maxConnsPerHost, 20, "http max connections per host")
	giconfig.Add(idleConnTimeout, 90*time.Second, "http idle connections timeout")
	giconfig.Add(disableKeepAlives, true, "http disable keep alives")
	giconfig.Add(forceHTTP2, true, "http force http2")
	giconfig.Add(tlsHandshakeTimeout, 10*time.Second, "TLS handshake timeout")
	giconfig.Add(timeout, 30*time.Second, "timeout")
	giconfig.Add(keepAlive, 15*time.Second, "keep alive")
	giconfig.Add(expectContinueTimeout, 1*time.Second, "expect continue timeout")
	giconfig.Add(dualStack, true, "dual stack")

}

func maxIdleConnValue() int {
	return giconfig.Int(maxIdleConn)
}

func maxIdleConnPerHostValue() int {
	return giconfig.Int(maxIdleConnPerHost)
}

func maxConnsPerHostValue() int {
	return giconfig.Int(maxConnsPerHost)
}

func idleConnTimeoutValue() time.Duration {
	return giconfig.Duration(idleConnTimeout)
}

func disableKeepAlivesValue() bool {
	return giconfig.Bool(disableKeepAlives)
}

func forceHTTP2Value() bool {
	return giconfig.Bool(forceHTTP2)
}

func tlsHandshakeTimeoutValue() time.Duration {
	return giconfig.Duration(tlsHandshakeTimeout)
}

func timeoutValue() time.Duration {
	return giconfig.Duration(timeout)
}

func keepAliveValue() time.Duration {
	return giconfig.Duration(keepAlive)
}

func expectContinueTimeoutValue() time.Duration {
	return giconfig.Duration(expectContinueTimeout)
}
