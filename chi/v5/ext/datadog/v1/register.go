package gichidatadog

import (
	"context"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gilog "github.com/b2wdigital/goignite/v2/log"
	chitrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/go-chi/chi"
)

func Register(ctx context.Context) (*gichi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)
	logger.Trace("enabling datadog middleware in chi")

	return &gichi.Config{
		Middlewares: []func(http.Handler) http.Handler{
			chitrace.Middleware(),
		},
	}, nil
}
