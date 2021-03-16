package gichihealth

import (
	"context"
	"encoding/json"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gilog "github.com/b2wdigital/goignite/v2/log"
	girestresponse "github.com/b2wdigital/goignite/v2/rest/response"
)

func Register(ctx context.Context) (*gichi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)

	healthRoute := getRoute()

	logger.Tracef("configuring health router on %s in chi", healthRoute)

	healthHandler := NewHealthHandler()

	return &gichi.Config{
		Routes: []gichi.ConfigRouter{
			{
				Method:      http.MethodGet,
				HandlerFunc: healthHandler.Get(ctx),
				Pattern:     healthRoute,
			},
		},
	}, nil
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(ctx context.Context) http.HandlerFunc {
	resp, httpCode := girestresponse.NewHealth(ctx)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(resp)
	}
}
