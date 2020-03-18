package elasticsearch

import (
	"time"

	"github.com/b2wdigital/goignite/pkg/config"

	"log"
)

const (
	Addresses             = "transport.client.elasticsearch.addresses"
	Username              = "transport.client.elasticsearch.username"
	Password              = "transport.client.elasticsearch.password"
	CloudID               = "transport.client.elasticsearch.cloudid"
	APIKey                = "transport.client.elasticsearch.apikey"
	CACert                = "transport.client.elasticsearch.cacert"
	RetryOnStatus         = "transport.client.elasticsearch.retryonstatus"
	DisableRetry          = "transport.client.elasticsearch.disableretry"
	EnableRetryOnTimeout  = "transport.client.elasticsearch.enableretryontimeout"
	MaxRetries            = "transport.client.elasticsearch.maxretries"
	DiscoverNodesOnStart  = "transport.client.elasticsearch.discovernodesonstart"
	DiscoverNodesInterval = "transport.client.elasticsearch.discovernodesinterval"
	EnableMetrics         = "transport.client.elasticsearch.enablemetrics"
	EnableDebugLogger     = "transport.client.elasticsearch.enabledebuglogger"
	RetryBackoff          = "transport.client.elasticsearch.retrybackoff"
	HealthEnabled         = "transport.client.elasticsearch.health.enabled"
	HealthDescription     = "transport.client.elasticsearch.health.description"
	HealthRequired        = "transport.client.elasticsearch.health.required"
)

func init() {

	log.Println("getting configurations for elastic search")

	config.Add(Addresses, []string{"http://127.0.0.1:9200"}, "a list of Elasticsearch nodes to use")
	config.Add(Username, "", "username for HTTP Basic Authentication")
	config.Add(Password, "", "password for HTTP Basic Authentication")
	config.Add(CloudID, "", "endpoint for the Elastic Service (https://elastic.co/cloud)")
	config.Add(APIKey, "", "base64-encoded token for authorization; if set, overrides username and password")
	config.Add(CACert, "", "PEM-encoded certificate authorities")
	config.Add(RetryOnStatus, []string{"502", "503", "504"}, "List of status codes for retry")
	config.Add(DisableRetry, false, "")
	config.Add(EnableRetryOnTimeout, false, "")
	config.Add(MaxRetries, 3, "")
	config.Add(DiscoverNodesOnStart, false, "discover nodes when initializing the client")
	config.Add(DiscoverNodesInterval, 0*time.Millisecond, "discover nodes periodically. Default: 0 (disabled)")
	config.Add(EnableMetrics, false, "enable the metrics collection")
	config.Add(EnableDebugLogger, false, "enable the debug logging")
	config.Add(RetryBackoff, 5*time.Millisecond, "optional backoff duration")
	config.Add(HealthEnabled, true, "enabled/disable health check")
	config.Add(HealthDescription, "default connection", "define health description")
	config.Add(HealthRequired, true, "define health description")

}
