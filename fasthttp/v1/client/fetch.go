package gifasthttp

import (
	"context"
	neturl "net/url"

	"github.com/valyala/fasthttp"
)

type Fetch struct {
	client              *fasthttp.Client
	beforeRequest       []func(ctx context.Context, o FetchOptions) context.Context
	afterRequest        []func(ctx context.Context, o FetchOptions, r Response)
	InterceptorResponse InterceptorResponse
}

func NewDefaultFetch(ctx context.Context) *Fetch {
	return &Fetch{
		client: NewDefaultClient(ctx),
	}
}

func NewFetch(ctx context.Context, options *Options) *Fetch {
	return &Fetch{
		client: NewClient(ctx, options),
	}
}

func (c *Fetch) Use(m Middleware) {
	c.beforeRequest = append(c.beforeRequest, m.OnBeforeRequest)
	c.afterRequest = append(c.afterRequest, m.OnAfterRequest)
}

func (c *Fetch) OnBeforeRequest(fn func(ctx context.Context, o FetchOptions) context.Context) {
	c.beforeRequest = append(c.beforeRequest, fn)
}

func (c *Fetch) OnAfterRequest(fn func(ctx context.Context, o FetchOptions, r Response)) {
	c.afterRequest = append(c.afterRequest, fn)
}

func (c *Fetch) Fetch(o FetchOptions) Response {

	url, e := neturl.Parse(o.Url)

	if e != nil {
		return Response{
			Error: e,
		}
	}

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)
	req.SetRequestURI(o.Url)

	ctx := o.Ctx
	for _, m := range c.beforeRequest {
		ctx = m(ctx, o)
	}

	if o.Header != nil {
		o.Header.VisitAll(func(key, value []byte) {
			req.Header.Set(string(key), string(value))
		})
	}

	req.Header.SetMethod(o.Method)

	if len(o.Body) != 0 {
		req.SetBody(o.Body)
	}
	err := c.client.DoTimeout(req, res, o.Timeout)

	var resHeader fasthttp.ResponseHeader
	res.Header.CopyTo(&resHeader)

	var bd = make([]byte, len(res.Body()))
	copy(bd, res.Body())

	response := Response{
		URL:        *url,
		Body:       bd,
		Header:     &resHeader,
		StatusCode: res.StatusCode(),
		Error:      err,
	}

	for _, f := range c.afterRequest {
		f(ctx, o, response)
	}

	if c.InterceptorResponse == nil {
		return response
	}

	return c.InterceptorResponse(response)
}
