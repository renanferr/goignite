package newrelic

import (
	"context"

	giants "github.com/b2wdigital/goignite/v2/ants/v2"
	"github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/newrelic/v3"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

type middleware struct {
}

func (i *middleware) Before(ctx context.Context) context.Context {

	if IsEnabled() || !newrelic.IsEnabled() {
		return ctx
	}

	logger := log.FromContext(ctx).WithTypeOf(*i)

	logger.Trace("creating go routine for newrelic")

	txn := newrelic.FromContext(ctx).NewGoroutine()

	logger.Debug("goroutine for newrelic successfully created in context")

	return nr.NewContext(ctx, txn)
}

func (i *middleware) After(ctx context.Context) {

	if IsEnabled() || !newrelic.IsEnabled() {
		return
	}

}

func NewMiddleware() giants.Middleware {
	log.Trace("creating newrelic middleware for ants")
	return &middleware{}
}
