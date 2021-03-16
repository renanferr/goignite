package gidatadog

import (
	"math"
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root                  = "gi.datadog"
	service               = root + ".service"
	env                   = root + ".env"
	enabled               = root + ".enabled"
	tags                  = root + ".tags"
	host                  = root + ".host"
	port                  = root + ".port"
	lambdaMode            = root + ".lambdaMode"
	analytics             = root + ".analytics"
	analyticsRate         = root + ".analyticsRate"
	debugMode             = root + ".debugMode"
	debugStack            = root + ".debugStack"
	httpClientRoot        = root + ".httpClient"
	maxIdleConnPerHost    = httpClientRoot + ".maxIdleConnPerHost"
	maxIdleConn           = httpClientRoot + ".maxIdleConn"
	maxConnsPerHost       = httpClientRoot + ".maxConnsPerHost"
	idleConnTimeout       = httpClientRoot + ".idleConnTimeout"
	disableKeepAlives     = httpClientRoot + ".disableKeepAlives"
	forceHTTP2            = httpClientRoot + ".forceHTTP2"
	tlsHandshakeTimeout   = httpClientRoot + ".TLSHandshakeTimeout"
	timeout               = httpClientRoot + ".timeout"
	keepAlive             = httpClientRoot + ".keepAlive"
	expectContinueTimeout = httpClientRoot + ".expectContinueTimeout"
	dualStack             = httpClientRoot + ".dualStack"
	logRoot               = root + ".log"
	logLevel              = logRoot + ".level"
)

func init() {

	giconfig.Add(service, "", "service name for datadog")
	giconfig.Add(env, "", "service env")
	giconfig.Add(enabled, true, "enables datadog")
	giconfig.Add(logLevel, "DEBUG", "log level")
	giconfig.Add(tags, map[string]string{}, "sets a key/value pair which will be set as a tag on all spans created by tracer. This option may be used multiple times")
	giconfig.Add(host, "localhost", "sets the address where the agent is located")
	giconfig.Add(port, "8126", "sets the port where the agent is located")
	giconfig.Add(lambdaMode, false, "enables lambda mode on the tracer, for use with AWS Lambda")
	giconfig.Add(analytics, false, "allows specifying whether Trace Search & Analytics should be enabled for integrations")
	giconfig.Add(analyticsRate, math.NaN(), "sets the global sampling rate for sampling APM events")
	giconfig.Add(debugMode, false, "enables debug mode on the tracer, resulting in more verbose logging.")
	giconfig.Add(debugStack, true, "can be used to globally enable or disable the collection of stack traces when spans finish with errors.")
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

func IsEnabled() bool {
	return giconfig.Bool(enabled)
}
