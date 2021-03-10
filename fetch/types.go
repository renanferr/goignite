package fetch

import (
	"context"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type Options struct {
	Url     string
	Method  string
	Header  *fasthttp.RequestHeader
	Timeout time.Duration
	Body    []byte
	Ctx     context.Context
}

type Fetch struct {
	client              *fasthttp.Client
	udBeforeRequest     []RequestMiddleware
	InterceptorResponse InterceptorResponse
	udAfterRequest      []ResponseMiddleware
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
	RequestMiddleware   (func(Options, context.Context) context.Context)
	ResponseMiddleware  (func(Options, Response, context.Context))
)
