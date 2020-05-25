package gihttp

import giconfig "github.com/b2wdigital/goignite/config"

const (
	ServerAddress     = "gi.http.server.addr"
	MaxHeaderBytes    = "gi.http.server.max-header-bytes"
	ReadHeaderTimeout = "gi.http.server.read-header-timeout"
	ReadTimeout       = "gi.http.server.read-timeout"
	WriteTimeout      = "gi.http.server.write-timeout"
	IdleTimeout       = "gi.http.server.idle-timeout"
)

func init() {
	giconfig.Add(ServerAddress, ":8080", "server address")
	giconfig.Add(MaxHeaderBytes, 1048576, "max header timeout")
	giconfig.Add(ReadHeaderTimeout, "1s", "read header timeout")
	giconfig.Add(ReadTimeout, "1", "read timeout")
	giconfig.Add(WriteTimeout, "7s", "write timeout ")
	giconfig.Add(IdleTimeout, "10m", "interval between question creation")
}
