package girestylogger

import (
	"context"
	"encoding/json"
	"time"

	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-resty/resty/v2"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("enabling logger middleware in resty")

	client.OnBeforeRequest(logBeforeResponse)
	client.OnAfterResponse(logAfterResponse)

	logger.Debug("logger middleware successfully enabled in resty")

	return nil
}

func logBeforeResponse(client *resty.Client, request *resty.Request) error {

	logger := gilog.FromContext(request.Context())

	requestHeaders, _ := json.Marshal(request.Header)

	requestBody, _ := json.Marshal(request.Body)

	logger = logger.
		WithFields(
			gilog.Fields{
				"rest_client_host":     client.HostURL,
				"rest_request_body":    string(requestBody),
				"rest_request_url":     request.URL,
				"rest_request_headers": string(requestHeaders),
				"rest_request_method":  request.Method,
			})

	var method func(format string, args ...interface{})

	switch Level() {
	case "TRACE":
		method = logger.Tracef
	case "DEBUG":
		method = logger.Debugf
	default:
		method = logger.Infof
	}

	method("rest request processing")

	return nil
}

func logAfterResponse(client *resty.Client, response *resty.Response) error {

	logger := gilog.FromContext(response.Request.Context())

	responseHeaders, _ := json.Marshal(response.Header())

	statusCode := response.StatusCode()

	logger = logger.WithFields(
		gilog.Fields{
			"rest_response_body":        string(response.Body()),
			"rest_response_headers":     string(responseHeaders),
			"rest_response_time":        response.Time().Seconds() * float64(time.Second/time.Millisecond),
			"rest_response_status_code": statusCode,
		})

	var method func(format string, args ...interface{})

	switch Level() {
	case "TRACE":
		method = logger.Tracef
	case "DEBUG":
		method = logger.Debugf
	default:
		method = logger.Infof
	}

	if statusCode > 500 {
		logger.Errorf("rest request processed with error")
	} else if statusCode > 400 {
		logger.Warnf("rest request processed with warning")
	} else {
		method("successful rest request processed")
	}

	return nil
}
