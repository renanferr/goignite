package gifasthttp

import (
	"context"
	"net/url"

	"github.com/valyala/fasthttp"
)

type Middleware interface {
	OnBeforeRequest(context.Context, FetchOptions) context.Context
	OnAfterRequest(context.Context, FetchOptions, Response)
}

type Response struct {
	URL        url.URL
	Body       []byte
	Header     *fasthttp.ResponseHeader
	StatusCode int
	Error      error
}

type (
	InterceptorResponse (func(Response) Response)
)
