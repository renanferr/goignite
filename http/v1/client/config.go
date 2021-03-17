package client

import (
	"time"

	"github.com/b2wdigital/goignite/v2/config"
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

	config.Add(maxIdleConnPerHost, 1, "http max idle connections per host")
	config.Add(maxIdleConn, 100, "http max idle connections")
	config.Add(maxConnsPerHost, 20, "http max connections per host")
	config.Add(idleConnTimeout, 90*time.Second, "http idle connections timeout")
	config.Add(disableKeepAlives, true, "http disable keep alives")
	config.Add(forceHTTP2, true, "http force http2")
	config.Add(tlsHandshakeTimeout, 10*time.Second, "TLS handshake timeout")
	config.Add(timeout, 30*time.Second, "timeout")
	config.Add(keepAlive, 15*time.Second, "keep alive")
	config.Add(expectContinueTimeout, 1*time.Second, "expect continue timeout")
	config.Add(dualStack, true, "dual stack")

}

func maxIdleConnValue() int {
	return config.Int(maxIdleConn)
}

func maxIdleConnPerHostValue() int {
	return config.Int(maxIdleConnPerHost)
}

func maxConnsPerHostValue() int {
	return config.Int(maxConnsPerHost)
}

func idleConnTimeoutValue() time.Duration {
	return config.Duration(idleConnTimeout)
}

func disableKeepAlivesValue() bool {
	return config.Bool(disableKeepAlives)
}

func forceHTTP2Value() bool {
	return config.Bool(forceHTTP2)
}

func tlsHandshakeTimeoutValue() time.Duration {
	return config.Duration(tlsHandshakeTimeout)
}

func timeoutValue() time.Duration {
	return config.Duration(timeout)
}

func keepAliveValue() time.Duration {
	return config.Duration(keepAlive)
}

func expectContinueTimeoutValue() time.Duration {
	return config.Duration(expectContinueTimeout)
}
