package fasthttp

import (
	"github.com/valyala/fasthttp"
)

func NewClient(o *Options) *fasthttp.Client {

	client := &fasthttp.Client{
		Name:                          o.Name,
		NoDefaultUserAgentHeader:      o.NoDefaultUserAgentHeader, // Don't send: User-Agent: fasthttp
		DialDualStack:                 o.DialDualStack,
		MaxConnsPerHost:               o.MaxConnsPerHost,
		MaxIdleConnDuration:           o.MaxIdleConnDuration,
		MaxIdemponentCallAttempts:     o.MaxIdemponentCallAttempts,
		ReadBufferSize:                o.ReadBufferSize, // Make sure to set this big enough that your whole request can be read at once.
		WriteBufferSize:               o.WriteBufferSize, // Same but for your response.
		ReadTimeout:                   o.ReadTimeout,
		WriteTimeout:                  o.WriteTimeout,
		MaxResponseBodySize:           o.MaxResponseBodySize,
		DisableHeaderNamesNormalizing: o.DisableHeaderNamesNormalizing, // If you set the case on your headers correctly you can enable this.
	}

	return client
}
