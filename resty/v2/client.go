package giresty

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-resty/resty/v2"
)

const (
	TopicClient = "topic:resty:client"
)

func NewClient(ctx context.Context, options *Options) *resty.Client {

	l := gilog.FromContext(ctx)

	l.Infof("creating resty client")

	client := resty.New()

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	transport := &http.Transport{
		DisableCompression:    giconfig.Bool(TransportDisableCompression),
		DisableKeepAlives:     giconfig.Bool(TransportDisableKeepAlives),
		MaxIdleConnsPerHost:   giconfig.Int(TransportMaxConnsPerHost),
		ResponseHeaderTimeout: giconfig.Duration(TransportResponseHeaderTimeout),
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     giconfig.Bool(TransportForceAttemptHTTP2),
		MaxIdleConns:          giconfig.Int(TransportMaxIdleConns),
		MaxConnsPerHost:       giconfig.Int(TransportMaxConnsPerHost),
		IdleConnTimeout:       giconfig.Duration(TransportIdleConnTimeout),
		TLSHandshakeTimeout:   giconfig.Duration(TransportTLSHandshakeTimeout),
		ExpectContinueTimeout: giconfig.Duration(TransportExpectContinueTimeout),
	}

	client.
		SetTransport(transport).
		SetTimeout(giconfig.Duration(RequestTimeout)).
		SetRetryCount(giconfig.Int(RetryCount)).
		SetRetryWaitTime(giconfig.Duration(RetryWaitTime)).
		SetRetryMaxWaitTime(giconfig.Duration(RetryMaxWaitTime)).
		SetDebug(false).
		SetHostURL(options.Host).
		AddRetryCondition(statusCodeRetryCondition)

	addTimeoutRetryCondition(client, options)

	if options.Debug || giconfig.Bool(Debug) {
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

	gieventbus.Publish(TopicClient, client)

	return client
}

func addTimeoutRetryCondition(client *resty.Client, options *Options) {

	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {

			var timeout time.Duration

			if options.RequestTimeout > 0 {
				timeout = options.RequestTimeout
			} else {
				timeout = giconfig.Duration(RequestTimeout)
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

	l := gilog.FromContext(request.Context())

	requestHeaders, _ := json.Marshal(request.Header)

	requestBody, _ := json.Marshal(request.Body)

	l = l.
		WithFields(
			gilog.Fields{
				"rest_request_body":    string(requestBody),
				"rest_request_url":     request.URL,
				"rest_request_headers": string(requestHeaders),
				"rest_request_method":  request.Method,
			})

	l.Debugf("rest request processing")

	return nil
}

func logAfterResponse(client *resty.Client, response *resty.Response) error {

	l := gilog.FromContext(response.Request.Context())

	responseHeaders, _ := json.Marshal(response.Header())

	statusCode := response.StatusCode()

	l = l.WithFields(
		gilog.Fields{
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
