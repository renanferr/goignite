package gihttp

import (
	"net/http"
)

// New returns a pointer with new Server
func NewDefault(handler http.Handler) *http.Server {
	opt, err := DefaultOptions()
	if err != nil {
		panic(err)
	}
	return New(handler, opt)
}

// New returns a pointer with new Server
func New(handler http.Handler, options *Options) *http.Server {
	return &http.Server{
		Addr:              options.Addr,
		Handler:           handler,
		MaxHeaderBytes:    options.MaxHeaderBytes,
		ReadTimeout:       options.ReadTimeout,
		ReadHeaderTimeout: options.ReadHeaderTimeout,
		WriteTimeout:      options.WriteTimeout,
		IdleTimeout:       options.IdleTimeout,
	}
}
