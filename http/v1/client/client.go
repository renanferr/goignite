package gihttp

import (
	"context"
	"net"
	"net/http"

	gilog "github.com/b2wdigital/goignite/v2/log"
)

func NewClient(ctx context.Context, options *Options) *http.Client {

	tr := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   options.Timeout,
			KeepAlive: options.KeepAlive,
			DualStack: options.DualStack,
		}).DialContext,
		TLSHandshakeTimeout:   options.TLSHandshakeTimeout,
		DisableKeepAlives:     options.DisableKeepAlives,
		MaxIdleConns:          options.MaxIdleConn,
		MaxIdleConnsPerHost:   options.MaxIdleConnPerHost,
		MaxConnsPerHost:       options.MaxConnsPerHost,
		IdleConnTimeout:       options.IdleConnTimeout,
		ForceAttemptHTTP2:     options.ForceHTTP2,
		ExpectContinueTimeout: options.ExpectContinueTimeout,
	}

	return &http.Client{Transport: tr}
}

func NewDefaultClient(ctx context.Context) *http.Client {

	logger := gilog.FromContext(ctx)

	o, err := DefaultOptions()
	if err != nil {
		logger.Fatalf(err.Error())
	}

	return NewClient(ctx, o)
}
