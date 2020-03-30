package http

import (
	"github.com/b2wdigital/goignite/pkg/config"
	"net/http"
)

// NewServer returns a pointer with new Server
func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              config.String(ServerAdress),
		Handler:           handler,
		MaxHeaderBytes:    config.Int(MaxHeaderBytes),
		ReadTimeout:       config.Duration(ReadTimeout),
		ReadHeaderTimeout: config.Duration(ReadHeaderTimeout),
		WriteTimeout:      config.Duration(WriteTimeout),
		IdleTimeout:       config.Duration(IdleTimeout),
	}
}
