package resty

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/http/client/resty/config"
	h "github.com/b2wdigital/goignite/pkg/http/client/resty/health"
	"github.com/b2wdigital/goignite/pkg/http/client/resty/model"
	"github.com/b2wdigital/goignite/pkg/log/logrus"

	c "github.com/b2wdigital/goignite/pkg/config"
)

func NewClient(ctx context.Context, options *model.Options) *resty.Client {

	log := logrus.FromContext(ctx)

	log.Info("creating resty client")

	client := resty.New()

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	transport := &http.Transport{
		DisableCompression:    c.Bool(config.TransportDisableCompression),
		DisableKeepAlives:     c.Bool(config.TransportDisableKeepAlives),
		MaxIdleConnsPerHost:   c.Int(config.TransportMaxConnsPerHost),
		ResponseHeaderTimeout: c.Duration(config.TransportResponseHeaderTimeout),
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     c.Bool(config.TransportForceAttemptHTTP2),
		MaxIdleConns:          c.Int(config.TransportMaxIdleConns),
		MaxConnsPerHost:       c.Int(config.TransportMaxConnsPerHost),
		IdleConnTimeout:       c.Duration(config.TransportIdleConnTimeout),
		TLSHandshakeTimeout:   c.Duration(config.TransportTLSHandshakeTimeout),
		ExpectContinueTimeout: c.Duration(config.TransportExpectContinueTimeout),
	}

	client.
		SetTransport(transport).
		SetTimeout(c.Duration(config.RequestTimeout)).
		SetRetryCount(c.Int(config.RetryCount)).
		SetRetryWaitTime(c.Duration(config.RetryWaitTime)).
		SetRetryMaxWaitTime(c.Duration(config.RetryMaxWaitTime)).
		SetDebug(false).
		SetHostURL(options.Host).
		AddRetryCondition(statusCodeRetryCondition)

	addTimeoutRetryCondition(client, options)


	if options.Debug || c.Bool(config.Debug) {
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

func addTimeoutRetryCondition(client *resty.Client, options *model.Options) {

	client.AddRetryCondition(
		func(r *resty.Response, err error) (bool, ) {

			var timeout time.Duration

			if options.RequestTimeout > 0 {
				timeout = options.RequestTimeout
			} else {
				timeout = c.Duration(config.RequestTimeout)
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

	log := logrus.FromContext(request.Context())

	requestHeaders, _ := json.Marshal(request.Header)

	requestBody, _ := json.Marshal(request.Body)

	logRest := log.
		WithField("rest_request_body", string(requestBody)).
		WithField("rest_request_url", request.URL).
		WithField("rest_request_headers", string(requestHeaders)).
		WithField("rest_request_method", request.Method)

	logRest.Trace("rest request processing")

	return nil
}

func logAfterResponse(client *resty.Client, response *resty.Response) error {

	log := logrus.FromContext(response.Request.Context())

	responseHeaders, _ := json.Marshal(response.Header())

	statusCode := response.StatusCode()

	logRest := log.
		WithField("rest_response_body", string(response.Body())).
		WithField("rest_response_headers", string(responseHeaders)).
		WithField("rest_response_time", response.Time().Seconds()*float64(time.Second/time.Millisecond)).
		WithField("rest_response_status_code", statusCode)

	if statusCode > 500 {
		logRest.Error("rest request processed with error")
	} else if statusCode > 400 {
		logRest.Warn("rest request processed with warning")
	} else {
		logRest.Debug("successful rest request processed")
	}

	return nil
}

func configureHealthCheck(client *resty.Client, o *model.Options) {

	mc := h.NewRestyChecker(client, o)
	hc := health.NewHealthChecker("http rest client", o.Health.Description, mc, o.Health.Required)

	health.Add(hc)
}
