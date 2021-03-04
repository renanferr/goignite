package gihttp

import (
	"time"

	giconfig "github.com/b2wdigital/goignite/v2/config"
)

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
	giconfig.Add(serverAddress, ":8081", "server address")
	giconfig.Add(maxHeaderBytes, 1048576, "max header timeout")
	giconfig.Add(readHeaderTimeout, 1*time.Second, "read header timeout")
	giconfig.Add(readTimeout, 1*time.Second, "read timeout")
	giconfig.Add(writeTimeout, 7*time.Second, "write timeout ")
	giconfig.Add(idleTimeout, 30*time.Second, "idle timeout")
}

func GetServerAddress() string {
	return giconfig.String(serverAddress)
}

func GetMaxHeaderBytes() int {
	return giconfig.Int(maxHeaderBytes)
}

func GetReadHeaderTimeout() time.Duration {
	return giconfig.Duration(readHeaderTimeout)
}

func GetReadTimeout() time.Duration {
	return giconfig.Duration(readTimeout)
}

func GetWriteTimeout() time.Duration {
	return giconfig.Duration(writeTimeout)
}

func GetIdleTimeout() time.Duration {
	return giconfig.Duration(idleTimeout)
}
