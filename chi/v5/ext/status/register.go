package status

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	gilog "github.com/b2wdigital/goignite/log"
	"github.com/b2wdigital/goignite/rest/response"
	"github.com/go-chi/chi/v5"
)

func Register(ctx context.Context, instance *chi.Mux) error {
	if !IsEnabled() {
		return nil
	}

	logger := gilog.FromContext(ctx)

	statusRoute := getRoute()

	logger.Infof("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()
	instance.Get(statusRoute, statusHandler.Get())

	return nil
}

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get() http.HandlerFunc {
	resourceStatus := response.NewResourceStatus()
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(resourceStatus)

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(reqBodyBytes.Bytes())
		w.WriteHeader(http.StatusOK)
	}
}
