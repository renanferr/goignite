package http

import "github.com/b2wdigital/goignite/pkg/config"

const (
	ServerAddress     = "transport.server.http.addr"
	MaxHeaderBytes    = "transport.server.http.max-header-bytes"
	ReadHeaderTimeout = "transport.server.http.read-header-timeout"
	ReadTimeout       = "transport.server.http.read-timeout"
	WriteTimeout      = "transport.server.http.write-timeout"
	IdleTimeout       = "transport.server.http.idle-timeout"
)

func init() {
	config.Add(ServerAddress, ":8080", "server address")
	config.Add(MaxHeaderBytes, 1048576, "max header timeout")
	config.Add(ReadHeaderTimeout, "1s", "read header timeout")
	config.Add(ReadTimeout, "1", "read timeout")
	config.Add(WriteTimeout, "7s", "write timeout ")
	config.Add(IdleTimeout, "10m", "interval between question creation")
}
