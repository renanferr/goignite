package retry

import (
	"context"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
	"github.com/go-resty/resty/v2"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !isEnabled() {
		return nil
	}

	client.
		SetRetryCount(giconfig.Int(count)).
		SetRetryWaitTime(giconfig.Duration(waitTime)).
		SetRetryMaxWaitTime(giconfig.Duration(maxWaitTime)).
		AddRetryCondition(statusCodeRetryCondition).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {

				timeout := client.GetClient().Timeout

				if r.Time() > timeout {
					return true
				}

				return false
			})

	addTimeoutRetryCondition(client)

	return nil
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

func addTimeoutRetryCondition(client *resty.Client) {

	client.AddRetryCondition(
		func(r *resty.Response, err error) bool {

			timeout := client.GetClient().Timeout

			if r.Time() > timeout {
				return true
			}

			return false
		})
}
