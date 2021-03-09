package giantsnewrelic

import (
	"context"

	giants "github.com/b2wdigital/goignite/v2/ants/v2"
	ginewrelic "github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type middleware struct {
}

func (i *middleware) Before(ctx context.Context) context.Context {
	txn := ginewrelic.FromContext(ctx).NewGoroutine()
	return newrelic.NewContext(ctx, txn)
}

func (i *middleware) After(ctx context.Context) {
	txn := ginewrelic.FromContext(ctx)
	if txn != nil && txn.IsSampled() {
		txn.End()
	}
}

func NewMiddleware() giants.Middleware {
	return &middleware{}
}
