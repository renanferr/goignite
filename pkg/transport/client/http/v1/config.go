package http

import (
	"log"
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Pkg                   = "pkg.provider.http"
	MaxIdleConnPerHost    = Pkg + ".maxidleconnperhost"
	MaxIdleConn           = Pkg + ".maxidleconn"
	MaxConnsPerHost       = Pkg + ".maxconnsperhost"
	IdleConnTimeout       = Pkg + ".idleconntimeout"
	DisableKeepAlives     = Pkg + ".disablekeepalives"
	ForceHTTP2            = Pkg + ".http2"
	TLSHandshakeTimeout   = Pkg + ".tlshandshaketimeout"
	Timeout               = Pkg + ".timeout"
	KeepAlive             = Pkg + ".keepalive"
	ExpectContinueTimeout = Pkg + ".expectcontinuetimeout"
	DualStack             = Pkg + ".dualstack"
)

func init() {

	log.Println("getting configurations for http client")

	config.Add(MaxIdleConnPerHost, 1, "http max idle connections per host")
	config.Add(MaxIdleConn, 100, "http max idle connections")
	config.Add(MaxConnsPerHost, 20, "http max connections per host")
	config.Add(IdleConnTimeout, 90*time.Second, "http idle connections timeout")
	config.Add(DisableKeepAlives, true, "http disable keep alives")
	config.Add(ForceHTTP2, true, "http force http2")
	config.Add(TLSHandshakeTimeout, 10*time.Second, "TLS handshake timeout")
	config.Add(Timeout, 30*time.Second, "timeout")
	config.Add(KeepAlive, 15*time.Second, "keep alive")
	config.Add(ExpectContinueTimeout, 1*time.Second, "expect continue timeout")
	config.Add(DualStack, true, "dual stack")

}

func MaxIdleConnValue() int {
	return config.Int(MaxIdleConn)
}

func MaxIdleConnPerHostValue() int {
	return config.Int(MaxIdleConnPerHost)
}

func MaxConnsPerHostValue() int {
	return config.Int(MaxConnsPerHost)
}

func IdleConnTimeoutValue() time.Duration {
	return config.Duration(IdleConnTimeout)
}

func DisableKeepAlivesValue() bool {
	return config.Bool(DisableKeepAlives)
}

func ForceHTTP2Value() bool {
	return config.Bool(ForceHTTP2)
}

func TLSHandshakeTimeoutValue() time.Duration {
	return config.Duration(TLSHandshakeTimeout)
}

func TimeoutValue() time.Duration {
	return config.Duration(Timeout)
}

func KeepAliveValue() time.Duration {
	return config.Duration(KeepAlive)
}

func ExpectContinueTimeoutValue() time.Duration {
	return config.Duration(ExpectContinueTimeout)
}
