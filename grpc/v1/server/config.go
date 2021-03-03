package gigrpc

import (
	giconfig "github.com/b2wdigital/goignite/v2/config"
)

const (
	root                 = "gi.grpc"
	port                 = root + ".port"
	maxConcurrentStreams = root + ".maxConcurrentStreams"
	tlsEnabled           = root + ".tls.enabled"
	certFile             = root + ".tls.certFile"
	keyFile              = root + ".tls.keyFile"
	caFile               = root + ".tls.CAFile"
	ExtRoot              = root + ".ext"
)

func init() {

	giconfig.Add(port, 9090, "server grpc port")
	giconfig.Add(maxConcurrentStreams, 5000, "server grpc max concurrent streams")
	giconfig.Add(tlsEnabled, false, "Use TLS - required for HTTP2.")
	giconfig.Add(certFile, "./cert/out/localhost.crt", "Path to the CRT/PEM file.")
	giconfig.Add(keyFile, "./cert/out/localhost.key", "Path to the private key file.")
	giconfig.Add(caFile, "./cert/out/blackbox.crt", "Path to the certificate authority (CA).")

}
