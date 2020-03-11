package resty

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"github.com/go-resty/resty/v2"
)

func NewClient(ctx context.Context, options *Options) *resty.Client {

	l := log.FromContext(ctx)

	l.Infof("creating resty client")

	client := resty.New()

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	transport := &http.Transport{
		DisableCompression:    config.Bool(TransportDisableCompression),
		DisableKeepAlives:     config.Bool(TransportDisableKeepAlives),
		MaxIdleConnsPerHost:   config.Int(TransportMaxConnsPerHost),
		ResponseHeaderTimeout: config.Duration(TransportResponseHeaderTimeout),
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     config.Bool(TransportForceAttemptHTTP2),
		MaxIdleConns:          config.Int(TransportMaxIdleConns),
		MaxConnsPerHost:       config.Int(TransportMaxConnsPerHost),
		IdleConnTimeout:       config.Duration(TransportIdleConnTimeout),
		TLSHandshakeTimeout:   config.Duration(TransportTLSHandshakeTimeout),
		ExpectContinueTimeout: config.Duration(TransportExpectContinueTimeout),
	}

	client.
		SetTransport(transport).
		SetTimeout(config.Duration(RequestTimeout)).
		SetRetryCount(config.Int(RetryCount)).
		SetRetryWaitTime(config.Duration(RetryWaitTime)).
		SetRetryMaxWaitTime(config.Duration(RetryMaxWaitTime)).
		SetDebug(false).
		SetHostURL(options.Host).
		AddRetryCondition(statusCodeRetryCondition)

	addTimeoutRetryCondition(client, options)

	if options.Debug || config.Bool(Debug) {
		client.OnBeforeRequest(logBeforeResponse)
		client.OnAfterResponse(logAfterResponse)
		client.SetDebug(true)
	}

	if options.RequestTimeout > -1 {
		client.SetTimeout(options.RequestTimeout)
	}

	if options.Retry.Count > -1 {
		client.SetRetryCount(options.Retry.Count)
	}

	if options.Retry.WaitTime > -1 {
		client.SetRetryWaitTime(options.Retry.WaitTime)
	}

	if options.Retry.MaxWaitTime > -1 {
		client.SetRetryMaxWaitTime(options.Retry.WaitTime)
	}

	if options.Health.Enabled {
		configureHealthCheck(client, options)
	}

	return client
}

func addTimeoutRetryCondition(client *resty.Client, options *Options) {

	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {

			var timeout time.Duration

			if options.RequestTimeout > 0 {
				timeout = options.RequestTimeout
			} else {
				timeout = config.Duration(RequestTimeout)
			}

			if r.Time() > timeout {
				return true
			}

			return false
		})
}

func statusCodeRetryCondition(r *resty.Response, err error) bool {
	switch statusCode := r.StatusCode(); statusCode {

	case http.StatusTooManyRequests:
		return true
	case http.StatusInternalServerError:
		return true
	case http.StatusGatewayTimeout:
		return true
	case http.StatusServiceUnavailable:
		return true
	default:
		return false
	}
}

func logBeforeResponse(client *resty.Client, request *resty.Request) error {

	l := log.FromContext(request.Context())

	requestHeaders, _ := json.Marshal(request.Header)

	requestBody, _ := json.Marshal(request.Body)

	l = l.
		WithFields(
			log.Fields{
				"rest_request_body":    string(requestBody),
				"rest_request_url":     request.URL,
				"rest_request_headers": string(requestHeaders),
				"rest_request_method":  request.Method,
			})

	l.Debugf("rest request processing")

	return nil
}

func logAfterResponse(client *resty.Client, response *resty.Response) error {

	l := log.FromContext(response.Request.Context())

	responseHeaders, _ := json.Marshal(response.Header())

	statusCode := response.StatusCode()

	l = l.WithFields(
		log.Fields{
			"rest_response_body":        string(response.Body()),
			"rest_response_headers":     string(responseHeaders),
			"rest_response_time":        response.Time().Seconds() * float64(time.Second/time.Millisecond),
			"rest_response_status_code": statusCode,
		})

	if statusCode > 500 {
		l.Errorf("rest request processed with error")
	} else if statusCode > 400 {
		l.Warnf("rest request processed with warning")
	} else {
		l.Debugf("successful rest request processed")
	}

	return nil
}

func configureHealthCheck(client *resty.Client, o *Options) {

	mc := NewRestyChecker(client, o)
	hc := health.NewHealthChecker("http rest client", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
