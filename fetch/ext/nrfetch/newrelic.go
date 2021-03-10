package nrfetch

import (
	"context"
	"net/http"

	"github.com/b2wdigital/goignite/v2/fetch"
	newrelic "github.com/newrelic/go-agent"
)

func Integrate(f *fetch.Fetch) {
	f.OnBeforeRequest(func(o fetch.Options, ctx context.Context) context.Context {
		reqHTTP, _ := http.NewRequest(o.Method, o.Url, nil)
		txn := newrelic.FromContext(ctx)
		seg := newrelic.StartExternalSegment(txn, reqHTTP)
		ctx = context.WithValue(ctx, "seg", seg)
		return ctx
	})

	f.OnAfterRequest(func(o fetch.Options, r fetch.Response, ctx context.Context) {
		seg := ctx.Value("seg").(*newrelic.ExternalSegment)
		seg.End()
	})
}
