package fetch

import (
	neturl "net/url"
	"time"

	"github.com/valyala/fasthttp"
)

func New() Fetch {
	client := fasthttp.Client{
		MaxConnsPerHost:     300,
		MaxConnDuration:     time.Duration(time.Second * 60),
		MaxIdleConnDuration: time.Duration(time.Second * 30),
	}

	f := Fetch{}

	f.client = &client

	return f
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
	for _, f := range c.udBeforeRequest {
		ctx = f(o, ctx)
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

	for _, f := range c.udAfterRequest {
		f(o, response, ctx)
	}

	if c.InterceptorResponse == nil {
		return response
	}

	return c.InterceptorResponse(response)
}

func (c *Fetch) OnBeforeRequest(m RequestMiddleware) {
	c.udBeforeRequest = append(c.udBeforeRequest, m)
}

func (c *Fetch) OnAfterRequest(m ResponseMiddleware) {
	c.udAfterRequest = append(c.udAfterRequest, m)
}
