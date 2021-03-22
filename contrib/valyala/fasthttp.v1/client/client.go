package client

import (
	"context"

	"github.com/b2wdigital/goignite/v2/core/log"
	"github.com/valyala/fasthttp"
)

func NewClient(ctx context.Context, o *Options) *fasthttp.Client {

	client := &fasthttp.Client{
		Name:                          o.Name,
		NoDefaultUserAgentHeader:      o.NoDefaultUserAgentHeader, // Don't send: User-Agent: fasthttp
		DialDualStack:                 o.DialDualStack,
		MaxConnsPerHost:               o.MaxConnsPerHost,
		MaxConnWaitTimeout:            o.MaxConnWaitTimeout,
		MaxIdleConnDuration:           o.MaxIdleConnDuration,
		MaxIdemponentCallAttempts:     o.MaxIdemponentCallAttempts,
		ReadBufferSize:                o.ReadBufferSize,  // Make sure to set this big enough that your whole request can be read at once.
		WriteBufferSize:               o.WriteBufferSize, // Same but for your response.
		ReadTimeout:                   o.ReadTimeout,
		WriteTimeout:                  o.WriteTimeout,
		MaxResponseBodySize:           o.MaxResponseBodySize,
		DisableHeaderNamesNormalizing: o.DisableHeaderNamesNormalizing, // If you set the case on your headers correctly you can enable this.
	}

	return client
}

func NewDefaultClient(ctx context.Context) *fasthttp.Client {

	logger := log.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}
