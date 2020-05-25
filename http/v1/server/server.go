package gihttp

import (
	"net/http"

	giconfig "github.com/b2wdigital/goignite/config"
)

// NewServer returns a pointer with new Server
func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              giconfig.String(ServerAddress),
		Handler:           handler,
		MaxHeaderBytes:    giconfig.Int(MaxHeaderBytes),
		ReadTimeout:       giconfig.Duration(ReadTimeout),
		ReadHeaderTimeout: giconfig.Duration(ReadHeaderTimeout),
		WriteTimeout:      giconfig.Duration(WriteTimeout),
		IdleTimeout:       giconfig.Duration(IdleTimeout),
	}
}
