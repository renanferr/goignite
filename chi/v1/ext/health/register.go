package health

import (
	"context"
	"encoding/json"
	"net/http"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/b2wdigital/goignite/rest/response"
	"github.com/go-chi/chi"
)

func Register(ctx context.Context, instance *chi.Mux) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	healthRoute := GetRoute()

	logger.Infof("configuring health router on %s", healthRoute)

	healthHandler := NewHealthHandler()
	instance.Get(healthRoute, healthHandler.Get(ctx))

	return nil
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

type HealthHandler struct {
}

func (u *HealthHandler) Get(ctx context.Context) http.HandlerFunc {
	resp, httpCode := response.NewHealth(ctx)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		json.NewEncoder(w).Encode(resp)
	}
}
