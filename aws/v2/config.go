package giaws

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	key            = "aws.access.key.id"
	secret         = "aws.secret.access.key"
	region         = "aws.default.region"
	session        = "aws.session.token"
	customEndpoint = "aws.custom.endpoint"

	retryer             = "aws.custom.retryer"
	retryerMaxAttempts  = retryer + ".maxattempts"
	retryerHasRateLimit = retryer + ".hasratelimit"
)

func init() {

	giconfig.Add(key, "", "defines the aws key")
	giconfig.Add(secret, "", "defines the aws secret")
	giconfig.Add(region, "", "defines the aws region")
	giconfig.Add(session, "", "defines the aws session token")
	giconfig.Add(customEndpoint, false, "defines if should point to localhost")

	giconfig.Add(retryerMaxAttempts, 5, "defines max attempts for rate limit")
	giconfig.Add(retryerHasRateLimit, true, "defines if retryer has rate limit")
}
