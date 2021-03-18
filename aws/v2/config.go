package aws

import (
	"time"

	"github.com/b2wdigital/goignite/v2/config"
)

const (
	key                 = "aws.access.key.id"
	secret              = "aws.secret.access.key"
	region              = "aws.default.region"
	accountNumber       = "aws.default.accountNumber"
	session             = "aws.session.token"
	customEndpoint      = "aws.custom.endpoint"
	retryerRoot         = "aws.custom.retryer"
	retryerMaxAttempts  = retryerRoot + ".maxattempts"
	retryerHasRateLimit = retryerRoot + ".hasratelimit"
	ExtRoot             = "aws.ext"
	httpClientRoot      = "aws.http.client"

	maxIdleConnPerHost    = httpClientRoot + ".maxIdleConnPerHost"
	maxIdleConn           = httpClientRoot + ".maxIdleConn"
	maxConnsPerHost       = httpClientRoot + ".maxConnsPerHost"
	idleConnTimeout       = httpClientRoot + ".idleConnTimeout"
	disableKeepAlives     = httpClientRoot + ".disableKeepAlives"
	disableCompression    = httpClientRoot + ".disableCompression"
	forceHTTP2            = httpClientRoot + ".forceHTTP2"
	tlsHandshakeTimeout   = httpClientRoot + ".TLSHandshakeTimeout"
	timeout               = httpClientRoot + ".timeout"
	dialTimeout           = httpClientRoot + ".dialTimeout"
	keepAlive             = httpClientRoot + ".keepAlive"
	expectContinueTimeout = httpClientRoot + ".expectContinueTimeout"
	dualStack             = httpClientRoot + ".dualStack"
)

func init() {

	config.Add(key, "", "defines the aws key")
	config.Add(secret, "", "defines the aws secret")
	config.Add(region, "", "defines the aws region")
	config.Add(accountNumber, "", "defines the aws account number")
	config.Add(session, "", "defines the aws session token")
	config.Add(customEndpoint, false, "defines if should point to localhost")

	config.Add(retryerMaxAttempts, 5, "defines max attempts for rate limit")
	config.Add(retryerHasRateLimit, true, "defines if retryer has rate limit")

	config.Add(maxIdleConnPerHost, 1, "http max idle connections per host")
	config.Add(maxIdleConn, 100, "http max idle connections")
	config.Add(maxConnsPerHost, 20, "http max connections per host")
	config.Add(idleConnTimeout, 90*time.Second, "http idle connections timeout")
	config.Add(disableKeepAlives, true, "http disable keep alives")
	config.Add(disableCompression, false, "http disable keep alives")
	config.Add(forceHTTP2, true, "http force http2")
	config.Add(tlsHandshakeTimeout, 10*time.Second, "TLS handshake timeout")
	config.Add(timeout, 30*time.Second, "timeout")
	config.Add(dialTimeout, 5*time.Second, "dial timeout")
	config.Add(keepAlive, 15*time.Second, "keep alive")
	config.Add(expectContinueTimeout, 1*time.Second, "expect continue timeout")
	config.Add(dualStack, true, "dual stack")
}

func Region() string {
	return config.String(region)
}

func AccountNumber() string {
	return config.String(accountNumber)
}
