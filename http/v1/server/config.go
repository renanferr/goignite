package gihttp

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root              = "gi.http.server"
	serverAddress     = root + ".addr"
	maxHeaderBytes    = root + ".maxHeaderBytes"
	readHeaderTimeout = root + ".readHeaderTimeout"
	readTimeout       = root + ".readTimeout"
	writeTimeout      = root + ".writeTimeout"
	idleTimeout       = root + ".idleTimeout"
)

func init() {
	giconfig.Add(serverAddress, ":8081", "server address")
	giconfig.Add(maxHeaderBytes, 1048576, "max header timeout")
	giconfig.Add(readHeaderTimeout, 1*time.Second, "read header timeout")
	giconfig.Add(readTimeout, 1*time.Second, "read timeout")
	giconfig.Add(writeTimeout, 7*time.Second, "write timeout ")
	giconfig.Add(idleTimeout, 30*time.Second, "idle timeout")
}
