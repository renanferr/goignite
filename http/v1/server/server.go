package gihttp

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

// New returns a pointer with new Server
func New(handler http.Handler) *http.Server {
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
