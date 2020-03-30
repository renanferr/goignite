package http

import "github.com/b2wdigital/goignite/pkg/config"

const (
	ServerAdress      = "transport.client.server.addr"
	MaxHeaderBytes    = "transport.client.server.max-header-bytes"
	ReadHeaderTimeout = "transport.client.server.read-header-timeout"
	ReadTimeout       = "transport.client.server.read-timeout"
	WriteTimeout      = "transport.client.server.write-timeout"
	IdleTimeout       = "transport.client.server.idle-timeout"
)

func init() {
	config.Add(ServerAdress, ":8080", "server address")
	config.Add(MaxHeaderBytes, 1048576, "max header timeout")
	config.Add(ReadHeaderTimeout, "1s", "read header timeout")
	config.Add(ReadTimeout, "1", "read timeout")
	config.Add(WriteTimeout, "7s", "write timeout ")
	config.Add(IdleTimeout, "10m", "interval between question creation")
}
