package gigrpc

import (
	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	root                 = "gi.grpc"
	Port                 = root + ".port"
	MaxConcurrentStreams = root + ".maxConcurrentStreams"
	TlsEnabled           = root + ".tls.enabled"
	CertFile             = root + ".tls.certFile"
	KeyFile              = root + ".tls.keyFile"
	CAFile               = root + ".tls.CAFile"
)

func init() {

	giconfig.Add(Port, 9090, "server grpc port")
	giconfig.Add(MaxConcurrentStreams, 5000, "server grpc max concurrent streams")
	giconfig.Add(TlsEnabled, false, "Use TLS - required for HTTP2.")
	giconfig.Add(CertFile, "./cert/out/localhost.crt", "Path to the CRT/PEM file.")
	giconfig.Add(KeyFile, "./cert/out/localhost.key", "Path to the private key file.")
	giconfig.Add(CAFile, "./cert/out/blackbox.crt", "Path to the certificate authority (CA).")

}
