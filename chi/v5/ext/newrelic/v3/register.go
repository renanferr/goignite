package gichinewrelic

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gilog "github.com/b2wdigital/goignite/v2/log"
	ginewrelic "github.com/b2wdigital/goignite/v2/newrelic/v3"
	"github.com/go-chi/chi/v5/middleware"
	nr "github.com/newrelic/go-agent/v3/newrelic"
)

func Register(ctx context.Context) (*gichi.Config, error) {
	if !IsEnabled() || !ginewrelic.IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("enabling newrelic middleware in chi")

	return &gichi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			nrMiddleware,
		},
	}, nil
}

func nrMiddleware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		url := r.URL.String()
		path := r.URL.Path
		txnName := strings.Join([]string{r.Method, path}, " ")

		txn := ginewrelic.Application().StartTransaction(txnName)
		defer txn.End()

		txn.SetWebRequestHTTP(r)

		if isWebResponseEnabled() {
			w = txn.SetWebResponse(w)
		}

		txn.AddAttribute("request.url", fmt.Sprintf("http://%s%s", r.Host, url))

		qs := r.URL.Query()
		for key, value := range qs {
			txn.AddAttribute(key, strings.Join(value, "|"))
		}

		if reqID := middleware.GetReqID(ctx); reqID != "" {
			txn.AddAttribute("request.id", reqID)
		}

		r = nr.RequestWithTransactionContext(r, txn)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
