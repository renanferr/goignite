package gifasthttpnewrelic

import (
	"context"
	"net/http"

	gifasthttp "github.com/b2wdigital/goignite/v2/fasthttp/v1/client"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const externalSegmentContextKey = "_fetch_newrelic_segment_"

type middleware struct {
}

func (m *middleware) OnBeforeRequest(ctx context.Context, o gifasthttp.FetchOptions) context.Context {
	reqHTTP, _ := http.NewRequest(o.Method, o.Url, nil)
	txn := newrelic.FromContext(ctx)
	s := newrelic.StartExternalSegment(txn, reqHTTP)
	ctx = context.WithValue(ctx, externalSegmentContextKey, s)
	return ctx
}

func (m *middleware) OnAfterRequest(ctx context.Context, o gifasthttp.FetchOptions, r gifasthttp.Response) {
	s := ctx.Value(externalSegmentContextKey).(*newrelic.ExternalSegment)
	s.End()
}

func New() gifasthttp.Middleware {
	return &middleware{}
}
