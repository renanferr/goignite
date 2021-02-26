package main

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (

	// config google client

	GoogleHost              = "app.client.resty.google.host"
	GoogleDebug             = "app.client.resty.google.debug"
	GoogleRequestTimeout    = "app.client.resty.google.request.timeout"
	GoogleRetryCount        = "app.client.resty.google.retry.count"
	GoogleRetryWaitTime     = "app.client.resty.google.retry.waittime"
	GoogleRetryMaxWaitTime  = "app.client.resty.google.retry.maxwaittime"
	GoogleHealthEnabled     = "app.client.resty.google.health.enabled"
	GoogleHealthDescription = "app.client.resty.google.health.description"
	GoogleHealthEndpoint    = "app.client.resty.google.health.endpoint"
	GoogleHealthRequired    = "app.client.resty.google.health.required"

	// config americanas client

	ACOMHost              = "app.client.resty.acom.host"
	ACOMDebug             = "app.client.resty.acom.debug"
	ACOMRequestTimeout    = "app.client.resty.acom.request.timeout"
	ACOMRetryCount        = "app.client.resty.acom.retry.count"
	ACOMRetryWaitTime     = "app.client.resty.acom.retry.waittime"
	ACOMRetryMaxWaitTime  = "app.client.resty.acom.retry.maxwaittime"
	ACOMHealthEnabled     = "app.client.resty.acom.health.enabled"
	ACOMHealthDescription = "app.client.resty.acom.health.description"
	ACOMHealthEndpoint    = "app.client.resty.acom.health.endpoint"
	ACOMHealthRequired    = "app.client.resty.acom.health.required"
)

func init() {

	giconfig.Add(GoogleHost, "http://www.google.com", "defines host")
	giconfig.Add(GoogleDebug, false, "defines client debug request")
	giconfig.Add(GoogleRequestTimeout, 2*time.Second, "defines client http request timeout (ms)")
	giconfig.Add(GoogleRetryCount, 0, "defines client max http retries")
	giconfig.Add(GoogleRetryWaitTime, 200*time.Millisecond, "defines client retry wait time (ms)")
	giconfig.Add(GoogleRetryMaxWaitTime, 2*time.Second, "defines client max retry wait time (ms)")
	giconfig.Add(GoogleHealthEnabled, true, "enable/disable health")
	giconfig.Add(GoogleHealthDescription, "google endpoint", "defines health description")
	giconfig.Add(GoogleHealthEndpoint, "http://www.google.com", "defines health endpoint")
	giconfig.Add(GoogleHealthRequired, true, "enable/disable health required dependency")

	giconfig.Add(ACOMHost, "http://www.americanas.com", "defines host")
	giconfig.Add(ACOMDebug, false, "defines client debug request")
	giconfig.Add(ACOMRequestTimeout, 2*time.Second, "defines client http request timeout (ms)")
	giconfig.Add(ACOMRetryCount, 0, "defines client max http retries")
	giconfig.Add(ACOMRetryWaitTime, 200*time.Millisecond, "defines client retry wait time (ms)")
	giconfig.Add(ACOMRetryMaxWaitTime, 2*time.Second, "defines client max retry wait time (ms)")
	giconfig.Add(ACOMHealthEnabled, true, "enable/disable health")
	giconfig.Add(ACOMHealthDescription, "google endpoint", "defines health description")
	giconfig.Add(ACOMHealthEndpoint, "http://www.google.com", "defines health endpoint")
	giconfig.Add(ACOMHealthRequired, true, "enable/disable health required dependency")

}
