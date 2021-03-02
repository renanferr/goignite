package logger

import (
	"context"
	"encoding/json"
	"time"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-resty/resty/v2"
)

func Middleware(ctx context.Context, client *resty.Client) error {

	if !isEnabled() {
		return nil
	}

	client.OnBeforeRequest(logBeforeResponse)
	client.OnAfterResponse(logAfterResponse)

	return nil
}

func logBeforeResponse(client *resty.Client, request *resty.Request) error {

	logger := gilog.FromContext(request.Context())

	requestHeaders, _ := json.Marshal(request.Header)

	requestBody, _ := json.Marshal(request.Body)

	logger = logger.
		WithFields(
			gilog.Fields{
				"rest_request_body":    string(requestBody),
				"rest_request_url":     request.URL,
				"rest_request_headers": string(requestHeaders),
				"rest_request_method":  request.Method,
			})

	logger.Infof("rest request processing")

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

	if statusCode > 500 {
		logger.Errorf("rest request processed with error")
	} else if statusCode > 400 {
		logger.Warnf("rest request processed with warning")
	} else {
		logger.Infof("successful rest request processed")
	}

	return nil
}
