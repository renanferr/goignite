package gifetch

import (
	neturl "net/url"
	"time"

	"github.com/valyala/fasthttp"
)

type Fetch struct {
	client              *fasthttp.Client
	middlewares         []Middleware
	InterceptorResponse InterceptorResponse
}

func New() Fetch {
	client := fasthttp.Client{
		MaxConnsPerHost:     300,
		MaxConnDuration:     60 * time.Second,
		MaxIdleConnDuration: 30 * time.Second,
	}

	f := Fetch{}

	f.client = &client

	return f
}

func (c *Fetch) Use(m Middleware) {
	c.middlewares = append(c.middlewares, m)
}

func (c *Fetch) Fetch(o Options) Response {

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
	for _, m := range c.middlewares {
		ctx = m.OnBeforeRequest(ctx, o)
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

	for _, f := range c.middlewares {
		f.OnAfterRequest(ctx, o, response)
	}

	if c.InterceptorResponse == nil {
		return response
	}

	return c.InterceptorResponse(response)
}
