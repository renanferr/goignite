package giresty

import (
	"context"

	"net/http"

	gieventbus "github.com/b2wdigital/goignite/eventbus"
	gilog "github.com/b2wdigital/goignite/log"
	giresty "github.com/b2wdigital/goignite/resty/v2"
	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type Integrator struct {
}

func Integrate() error {
	integrator := &Integrator{}
	return gieventbus.Subscribe(giresty.TopicClient, integrator.Integrate)
}

func (i *Integrator) Integrate(client *resty.Client) error {

	logger := gilog.WithTypeOf(*i)

	logger.Trace("integrating resty with newrelic")

	if IsEnabled() {

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
	} else {
		logger.Debug("resty integration is disabled")
	}

	return nil
}
