package giantsnewrelic

import (
	"context"

	giants "github.com/b2wdigital/goignite/v2/ants/v2"
	gilog "github.com/b2wdigital/goignite/v2/log"
	ginewrelic "github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type middleware struct {
}

func (i *middleware) Before(ctx context.Context) context.Context {

	if IsEnabled() || !ginewrelic.IsEnabled() {
		return ctx
	}

	logger := gilog.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("creating go routine for newrelic")

	txn := ginewrelic.FromContext(ctx).NewGoroutine()

	logger.Debug("goroutine for newrelic successfully created in context")

	return newrelic.NewContext(ctx, txn)
}

func (i *middleware) After(ctx context.Context) {

	if IsEnabled() || !ginewrelic.IsEnabled() {
		return
	}

}

func NewMiddleware() giants.Middleware {
	gilog.Trace("creating newrelic middleware for ants")
	return &middleware{}
}
