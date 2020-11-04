package gihttp

import (
	"log"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Pkg                   = "gi.http.client"
	MaxIdleConnPerHost    = Pkg + ".maxIdleConnPerHost"
	MaxIdleConn           = Pkg + ".maxIdleConn"
	MaxConnsPerHost       = Pkg + ".maxConnsPerHost"
	IdleConnTimeout       = Pkg + ".idleConnTimeout"
	DisableKeepAlives     = Pkg + ".disableKeepAlives"
	ForceHTTP2            = Pkg + ".forceHTTP2"
	TLSHandshakeTimeout   = Pkg + ".TLSHandshakeTimeout"
	Timeout               = Pkg + ".timeout"
	KeepAlive             = Pkg + ".keepAlive"
	ExpectContinueTimeout = Pkg + ".expectContinueTimeout"
	DualStack             = Pkg + ".dualStack"
)

func init() {

	log.Println("getting configurations for http client")

	giconfig.Add(MaxIdleConnPerHost, 1, "http max idle connections per host")
	giconfig.Add(MaxIdleConn, 100, "http max idle connections")
	giconfig.Add(MaxConnsPerHost, 20, "http max connections per host")
	giconfig.Add(IdleConnTimeout, 90*time.Second, "http idle connections timeout")
	giconfig.Add(DisableKeepAlives, true, "http disable keep alives")
	giconfig.Add(ForceHTTP2, true, "http force http2")
	giconfig.Add(TLSHandshakeTimeout, 10*time.Second, "TLS handshake timeout")
	giconfig.Add(Timeout, 30*time.Second, "timeout")
	giconfig.Add(KeepAlive, 15*time.Second, "keep alive")
	giconfig.Add(ExpectContinueTimeout, 1*time.Second, "expect continue timeout")
	giconfig.Add(DualStack, true, "dual stack")

}

func MaxIdleConnValue() int {
	return giconfig.Int(MaxIdleConn)
}

func MaxIdleConnPerHostValue() int {
	return giconfig.Int(MaxIdleConnPerHost)
}

func MaxConnsPerHostValue() int {
	return giconfig.Int(MaxConnsPerHost)
}

func IdleConnTimeoutValue() time.Duration {
	return giconfig.Duration(IdleConnTimeout)
}

func DisableKeepAlivesValue() bool {
	return giconfig.Bool(DisableKeepAlives)
}

func ForceHTTP2Value() bool {
	return giconfig.Bool(ForceHTTP2)
}

func TLSHandshakeTimeoutValue() time.Duration {
	return giconfig.Duration(TLSHandshakeTimeout)
}

func TimeoutValue() time.Duration {
	return giconfig.Duration(Timeout)
}

func KeepAliveValue() time.Duration {
	return giconfig.Duration(KeepAlive)
}

func ExpectContinueTimeoutValue() time.Duration {
	return giconfig.Duration(ExpectContinueTimeout)
}
