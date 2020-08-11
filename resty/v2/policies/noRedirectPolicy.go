package policies

import (
	"github.com/b2wdigital/goignite/errors"
	"net/http"
)

type NoRedirectPolicy struct {
}

func (p *NoRedirectPolicy) Apply(request *http.Request, requests []*http.Request) error {
	return errors.ForbiddenRedirectf("redirect is forbidden in config")
}
