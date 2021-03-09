package gielasticsearch

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root                  = "gi.elasticsearch"
	addresses             = root + ".addresses"
	username              = root + ".username"
	password              = root + ".password"
	cloudID               = root + ".cloudID"
	apiKey                = root + ".APIKey"
	caCert                = root + ".CACert"
	retryOnStatus         = root + ".retryOnStatus"
	disableRetry          = root + ".disableRetry"
	enableRetryOnTimeout  = root + ".enableRetryOnTimeout"
	maxRetries            = root + ".maxRetries"
	discoverNodesOnStart  = root + ".discoverNodesOnStart"
	discoverNodesInterval = root + ".discoverNodesInterval"
	enableMetrics         = root + ".enableMetrics"
	enableDebugLogger     = root + ".enableDebugLogger"
	retryBackoff          = root + ".retryBackoff"
	ExtRoot               = root + ".ext"
)

func init() {
	giconfig.Add(addresses, []string{"http://127.0.0.1:9200"}, "a list of Elasticsearch nodes to use")
	giconfig.Add(username, "", "username for HTTP Basic Authentication")
	giconfig.Add(password, "", "password for HTTP Basic Authentication")
	giconfig.Add(cloudID, "", "endpoint for the Elastic Service (https://elastic.co/cloud)")
	giconfig.Add(apiKey, "", "base64-encoded token for authorization; if set, overrides username and password")
	giconfig.Add(caCert, "", "PEM-encoded certificate authorities")
	giconfig.Add(retryOnStatus, []int{502, 503, 504}, "List of status codes for retry")
	giconfig.Add(disableRetry, false, "")
	giconfig.Add(enableRetryOnTimeout, false, "")
	giconfig.Add(maxRetries, 3, "")
	giconfig.Add(discoverNodesOnStart, false, "discover nodes when initializing the client")
	giconfig.Add(discoverNodesInterval, 0*time.Millisecond, "discover nodes periodically. Default: 0 (disabled)")
	giconfig.Add(enableMetrics, false, "enable the metrics collection")
	giconfig.Add(enableDebugLogger, false, "enable the debug logging")
	giconfig.Add(retryBackoff, 5*time.Millisecond, "optional backoff duration")
}
