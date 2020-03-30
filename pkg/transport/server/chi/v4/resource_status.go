package chi

import (
	"bytes"
	"encoding/json"
	"github.com/b2wdigital/goignite/pkg/rest/response"
	"net/http"
)

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
