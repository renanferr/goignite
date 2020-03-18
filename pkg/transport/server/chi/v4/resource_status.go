package chi

import (
	"encoding/json"
	"github.com/b2wdigital/goignite/pkg/info"
	"net/http"
)

func NewResourceStatusHandler() *ResourceStatusHandler {
	return &ResourceStatusHandler{}
}

type ResourceStatusHandler struct {
}

func (u *ResourceStatusHandler) Get() http.HandlerFunc {

	body,_:= json.Marshal(map[string]string{
		"applicationName": info.AppName,
	})

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
