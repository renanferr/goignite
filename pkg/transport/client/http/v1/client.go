package http

import (
	"net"
	"net/http"
)

func NewClient(options *Options) *http.Client {

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
		ForceAttemptHTTP2:     options.HTTP2,
		ExpectContinueTimeout: options.ExpectContinueTimeout,
	}
	return &http.Client{Transport: tr}

}
