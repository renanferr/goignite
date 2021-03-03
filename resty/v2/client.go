package giresty

import (
	"context"
	"net"
	"net/http"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	"github.com/go-resty/resty/v2"
)

type Ext func(context.Context, *resty.Client) error

func NewClient(ctx context.Context, options *Options, exts ...Ext) *resty.Client {

	logger := gilog.FromContext(ctx)

	logger.Infof("creating resty client")

	client := resty.New()

	dialer := &net.Dialer{
		Timeout:       giconfig.Duration(connectionTimeout),
		FallbackDelay: giconfig.Duration(fallbackDelay),
		KeepAlive:     giconfig.Duration(keepAlive),
	}

	if options.ConnectionTimeout > 0 {
		dialer.Timeout = options.ConnectionTimeout
	}

	if options.KeepAlive > 0 {
		dialer.KeepAlive = options.KeepAlive
	}

	transport := &http.Transport{
		DisableCompression:    giconfig.Bool(transportDisableCompression),
		DisableKeepAlives:     giconfig.Bool(transportDisableKeepAlives),
		MaxIdleConnsPerHost:   giconfig.Int(transportMaxConnsPerHost),
		ResponseHeaderTimeout: giconfig.Duration(transportResponseHeaderTimeout),
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     giconfig.Bool(transportForceAttemptHTTP2),
		MaxIdleConns:          giconfig.Int(transportMaxIdleConns),
		MaxConnsPerHost:       giconfig.Int(transportMaxConnsPerHost),
		IdleConnTimeout:       giconfig.Duration(transportIdleConnTimeout),
		TLSHandshakeTimeout:   giconfig.Duration(transportTLSHandshakeTimeout),
		ExpectContinueTimeout: giconfig.Duration(transportExpectContinueTimeout),
	}

	if options.Transport != nil {

		transport.DisableCompression = options.Transport.DisableCompression
		transport.DisableKeepAlives = options.Transport.DisableKeepAlives
		transport.ForceAttemptHTTP2 = options.Transport.ForceAttemptHTTP2

		if options.Transport.MaxIdleConnsPerHost > 0 {
			transport.MaxIdleConnsPerHost = options.Transport.MaxIdleConnsPerHost
		}

		if options.Transport.ResponseHeaderTimeout > 0 {
			transport.ResponseHeaderTimeout = options.Transport.ResponseHeaderTimeout
		}

		if options.Transport.MaxIdleConns > 0 {
			transport.MaxIdleConns = options.Transport.MaxIdleConns
		}

		if options.Transport.MaxConnsPerHost > 0 {
			transport.MaxConnsPerHost = options.Transport.MaxConnsPerHost
		}

		if options.Transport.IdleConnTimeout > 0 {
			transport.IdleConnTimeout = options.Transport.IdleConnTimeout
		}

		if options.Transport.TLSHandshakeTimeout > 0 {
			transport.TLSHandshakeTimeout = options.Transport.TLSHandshakeTimeout
		}

		if options.Transport.ExpectContinueTimeout > 0 {
			transport.ExpectContinueTimeout = options.Transport.ExpectContinueTimeout
		}
	}

	client.
		SetTransport(transport).
		SetTimeout(giconfig.Duration(requestTimeout)).
		SetDebug(giconfig.Bool(debug)).
		SetHostURL(options.Host).
		SetCloseConnection(giconfig.Bool(closeConnection))

	if options.Debug || giconfig.Bool(debug) {
		client.SetDebug(true)
	}

	if options.RequestTimeout > 0 {
		client.SetTimeout(options.RequestTimeout)
	}

	for _, ext := range exts {
		if err := ext(ctx, client); err != nil {
			panic(err)
		}
	}

	return client
}
