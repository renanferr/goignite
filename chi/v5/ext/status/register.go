package gichistatus

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	gichi "github.com/b2wdigital/goignite/v2/chi/v5"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/b2wdigital/goignite/v2/rest/response"
)

func Register(ctx context.Context) (*gichi.Config, error) {
	if !IsEnabled() {
		return nil, nil
	}

	logger := gilog.FromContext(ctx)

	statusRoute := getRoute()

	logger.Tracef("configuring status router on %s", statusRoute)

	statusHandler := NewResourceStatusHandler()

	return &gichi.Config{
		Routes: []gichi.ConfigRouter{
			{
				Method:      http.MethodGet,
				HandlerFunc: statusHandler.Get(),
				Pattern:     statusRoute,
			},
		},
	}, nil
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
