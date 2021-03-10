package gifetch

import (
	"context"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type Middleware interface {
	OnBeforeRequest(context.Context, Options) context.Context
	OnAfterRequest(context.Context, Options, Response)
}

type Options struct {
	Url     string
	Method  string
	Header  *fasthttp.RequestHeader
	Timeout time.Duration
	Body    []byte
	Ctx     context.Context
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
