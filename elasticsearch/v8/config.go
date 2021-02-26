package gielasticsearch

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Addresses             = "gi.elasticsearch.addresses"
	Username              = "gi.elasticsearch.username"
	Password              = "gi.elasticsearch.password"
	CloudID               = "gi.elasticsearch.cloudID"
	APIKey                = "gi.elasticsearch.APIKey"
	CACert                = "gi.elasticsearch.CACert"
	RetryOnStatus         = "gi.elasticsearch.retryOnStatus"
	DisableRetry          = "gi.elasticsearch.disableRetry"
	EnableRetryOnTimeout  = "gi.elasticsearch.enableRetryOnTimeout"
	MaxRetries            = "gi.elasticsearch.maxRetries"
	DiscoverNodesOnStart  = "gi.elasticsearch.discoverNodesOnStart"
	DiscoverNodesInterval = "gi.elasticsearch.discoverNodesInterval"
	EnableMetrics         = "gi.elasticsearch.enableMetrics"
	EnableDebugLogger     = "gi.elasticsearch.enableDebugLogger"
	RetryBackoff          = "gi.elasticsearch.retryBackoff"
)

func init() {

	giconfig.Add(Addresses, []string{"http://127.0.0.1:9200"}, "a list of Elasticsearch nodes to use")
	giconfig.Add(Username, "", "username for HTTP Basic Authentication")
	giconfig.Add(Password, "", "password for HTTP Basic Authentication")
	giconfig.Add(CloudID, "", "endpoint for the Elastic Service (https://elastic.co/cloud)")
	giconfig.Add(APIKey, "", "base64-encoded token for authorization; if set, overrides username and password")
	giconfig.Add(CACert, "", "PEM-encoded certificate authorities")
	giconfig.Add(RetryOnStatus, []int{502, 503, 504}, "List of status codes for retry")
	giconfig.Add(DisableRetry, false, "")
	giconfig.Add(EnableRetryOnTimeout, false, "")
	giconfig.Add(MaxRetries, 3, "")
	giconfig.Add(DiscoverNodesOnStart, false, "discover nodes when initializing the client")
	giconfig.Add(DiscoverNodesInterval, 0*time.Millisecond, "discover nodes periodically. Default: 0 (disabled)")
	giconfig.Add(EnableMetrics, false, "enable the metrics collection")
	giconfig.Add(EnableDebugLogger, false, "enable the debug logging")
	giconfig.Add(RetryBackoff, 5*time.Millisecond, "optional backoff duration")

}
