package gihttp

import giconfig "github.com/b2wdigital/goignite/config"

const (
	root              = "gi.http.server"
	serverAddress     = root + ".addr"
	maxHeaderBytes    = root + ".max-header-bytes"
	readHeaderTimeout = root + ".read-header-timeout"
	readTimeout       = root + ".read-timeout"
	writeTimeout      = root + ".write-timeout"
	idleTimeout       = root + ".idle-timeout"
)

func init() {
	giconfig.Add(serverAddress, ":8080", "server address")
	giconfig.Add(maxHeaderBytes, 1048576, "max header timeout")
	giconfig.Add(readHeaderTimeout, "1s", "read header timeout")
	giconfig.Add(readTimeout, "1", "read timeout")
	giconfig.Add(writeTimeout, "7s", "write timeout ")
	giconfig.Add(idleTimeout, "10m", "interval between question creation")
}
