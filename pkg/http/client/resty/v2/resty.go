package v1

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/go-resty/resty/v2"
    log "github.com/sirupsen/logrus"

    "github.com/jpfaria/goignite/pkg/http/client/resty/config"
    "github.com/jpfaria/goignite/pkg/http/client/resty/model"

    c "github.com/jpfaria/goignite/pkg/config"
)

func NewClient(options *model.Options) *resty.Client {

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

    if options.RetryCount > -1 {
        client.SetRetryCount(options.RetryCount)
    }

    if options.RetryWaitTime > -1 {
        client.SetRetryWaitTime(time.Duration(options.RetryWaitTime) * time.Millisecond)
    }

    if options.RetryMaxWaitTime > -1 {
        client.SetRetryMaxWaitTime(time.Duration(options.RetryMaxWaitTime) * time.Millisecond)
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
