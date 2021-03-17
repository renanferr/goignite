package aws

import (
	"time"

	"github.com/b2wdigital/goignite/v2/config"
)

const (
	key                         = "aws.access.key.id"
	secret                      = "aws.secret.access.key"
	region                      = "aws.default.region"
	accountNumber               = "aws.default.accountNumber"
	session                     = "aws.session.token"
	customEndpoint              = "aws.custom.endpoint"
	retryerRoot                 = "aws.custom.retryer"
	retryerMaxAttempts          = retryerRoot + ".maxattempts"
	retryerHasRateLimit         = retryerRoot + ".hasratelimit"
	ExtRoot                     = "aws.ext"
	httpClientRoot              = "aws.httpClient"
	maxConnsPerHost             = httpClientRoot + ".maxConnsPerHost"
	maxIdleConns                = httpClientRoot + ".maxIdleConns"
	maxIdleConnsPerHost         = httpClientRoot + ".maxIdleConnsPerHost"
	timeoutMillis               = httpClientRoot + ".timeoutMillis"
	keepAliveMillis             = httpClientRoot + ".keepAliveMillis"
	idleConnTimeoutMillis       = httpClientRoot + ".idleConnTimeoutMillis"
	responseHeaderTimeoutMillis = httpClientRoot + ".responseHeaderTimeoutMillis"
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

	config.Add(maxConnsPerHost, 256, "limits the total number of connections per host")
	config.Add(maxIdleConns, 100, "controls the maximum number of idle (keep-alive) connections across all hosts")
	config.Add(maxIdleConnsPerHost, 10, "controls the maximum idle (keep-alive) connections to keep per-host")
	config.Add(timeoutMillis, 10*time.Second, "the maximum amount of time a dial will wait for a connect to complete")
	config.Add(keepAliveMillis, 10*time.Second, "specifies the interval between keep-alive probes for an active network connection")
	config.Add(idleConnTimeoutMillis, 5*time.Second, "the maximum amount of time an idle (keep-alive) connection will remain idle before closing itself")
	config.Add(responseHeaderTimeoutMillis, 5*time.Second, "specifies the amount of time to wait for a server's response headers after fully writing the request (including its body, if any)")
}

func Region() string {
	return config.String(region)
}

func AccountNumber() string {
	return config.String(accountNumber)
}
