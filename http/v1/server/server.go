package gihttp

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
)

// NewServer returns a pointer with new Server
func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              giconfig.String(serverAddress),
		Handler:           handler,
		MaxHeaderBytes:    giconfig.Int(maxHeaderBytes),
		ReadTimeout:       giconfig.Duration(readTimeout),
		ReadHeaderTimeout: giconfig.Duration(readHeaderTimeout),
		WriteTimeout:      giconfig.Duration(writeTimeout),
		IdleTimeout:       giconfig.Duration(idleTimeout),
	}
}
