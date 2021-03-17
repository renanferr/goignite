package server

import "github.com/b2wdigital/goignite/v2/config"

const (
	root                 = "gi.grpc.server"
	port                 = root + ".port"
	maxConcurrentStreams = root + ".maxConcurrentStreams"
	tlsEnabled           = root + ".tls.enabled"
	certFile             = root + ".tls.certFile"
	keyFile              = root + ".tls.keyFile"
	caFile               = root + ".tls.CAFile"
	ExtRoot              = root + ".ext"
)

func init() {

	config.Add(port, 9090, "server grpc port")
	config.Add(maxConcurrentStreams, 5000, "server grpc max concurrent streams")
	config.Add(tlsEnabled, false, "Use TLS - required for HTTP2.")
	config.Add(certFile, "./cert/out/localhost.crt", "Path to the CRT/PEM file.")
	config.Add(keyFile, "./cert/out/localhost.key", "Path to the private key file.")
	config.Add(caFile, "./cert/out/blackbox.crt", "Path to the certificate authority (CA).")

}
