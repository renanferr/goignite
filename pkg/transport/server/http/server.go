package http

import (
	"net/http"

	"github.com/b2wdigital/goignite/pkg/config"
)

// NewServer returns a pointer with new Server
func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              config.String(ServerAddress),
		Handler:           handler,
		MaxHeaderBytes:    config.Int(MaxHeaderBytes),
		ReadTimeout:       config.Duration(ReadTimeout),
		ReadHeaderTimeout: config.Duration(ReadHeaderTimeout),
		WriteTimeout:      config.Duration(WriteTimeout),
		IdleTimeout:       config.Duration(IdleTimeout),
	}
}
