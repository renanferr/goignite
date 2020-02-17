package v2

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/jpfaria/goignite/pkg/health"
	"github.com/jpfaria/goignite/pkg/http/client/resty/config"
	h "github.com/jpfaria/goignite/pkg/http/client/resty/health"
	"github.com/jpfaria/goignite/pkg/http/client/resty/model"
	"github.com/jpfaria/goignite/pkg/logging/logrus"

	c "github.com/jpfaria/goignite/pkg/config"
)

func NewClient(ctx context.Context, options *model.Options) *resty.Client {

	log := logrus.FromContext(ctx)

	log.Trace("creating resty client")

	client := resty.New()

	client.
		SetTimeout(time.Duration(c.Int(config.RequestTimeout)) * time.Millisecond).
		SetRetryCount(c.Int(config.RetryCount)).
		SetRetryWaitTime(time.Duration(c.Int(config.RetryWaitTime)) * time.Millisecond).
		SetRetryMaxWaitTime(time.Duration(c.Int(config.RetryMaxWaitTime)) * time.Millisecond).
		SetDebug(false).
		SetHostURL(options.Host).
		AddRetryCondition(statusCodeRetryCondition).
		AddRetryCondition(
			func(r *resty.Response, err error) (bool, ) {

				if r.Time() > time.Duration(options.RequestTimeout)*time.Millisecond {
					return true
				}

				return false
			})

	if options.Debug || c.Bool(config.Debug) {
		client.OnBeforeRequest(logBeforeResponse)
		client.OnAfterResponse(logAfterResponse)
		client.SetDebug(true)
	}

	if options.RequestTimeout > -1 {
		client.SetTimeout(time.Duration(options.RequestTimeout) * time.Millisecond)
	}

	if options.Retry.Count > -1 {
		client.SetRetryCount(options.Retry.Count)
	}

	if options.Retry.WaitTime > -1 {
		client.SetRetryWaitTime(time.Duration(options.Retry.WaitTime) * time.Millisecond)
	}

	if options.Retry.MaxWaitTime > -1 {
		client.SetRetryMaxWaitTime(time.Duration(options.Retry.MaxWaitTime) * time.Millisecond)
	}

	if options.Health.Enabled {
		configureHealthCheck(client, options)
	}

	return client
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
