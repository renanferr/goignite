package giresty

import (
	"context"

	"net/http"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func Register(ctx context.Context, client *resty.Client) error {

	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	logger.Trace("integrating resty with newrelic")

	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {

		rctx := request.Context()

		logger := gilog.FromContext(rctx)

		txn := newrelic.FromContext(rctx)
		txn.InsertDistributedTraceHeaders(request.Header)

		req, _ := http.NewRequest(request.Method, client.HostURL, nil)
		req.Header = request.Header

		s := newrelic.StartExternalSegment(txn, req)
		ctx := context.WithValue(rctx, "nrext", s)

		request.SetContext(ctx)

		logger.Debugf("rest request processing")

		return nil
	})

	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {

		ctx := resp.Request.Context()

		s := ctx.Value("nrext").(*newrelic.ExternalSegment)
		s.End()

		return nil
	})

	logger.Debug("resty integrated with newrelic with success")

	return nil
}
