package gigrpc

import (
	"log"

	giconfig "github.com/b2wdigital/goignite/config"
)

const (
	Port                 = "gi.grpc.port"
	MaxConcurrentStreams = "gi.grpc.maxConcurrentStreams"
	TlsEnabled           = "gi.grpc.tls.enabled"
	CertFile             = "gi.grpc.tls.certFile"
	KeyFile              = "gi.grpc.tls.keyFile"
	CAFile               = "gi.grpc.tls.CAFile"
)

func init() {

	log.Println("getting configurations for grpc server")

	giconfig.Add(Port, 9090, "server grpc port")
	giconfig.Add(MaxConcurrentStreams, 5000, "server grpc max concurrent streams")
	giconfig.Add(TlsEnabled, false, "Use TLS - required for HTTP2.")
	giconfig.Add(CertFile, "./cert/out/localhost.crt", "Path to the CRT/PEM file.")
	giconfig.Add(KeyFile, "./cert/out/localhost.key", "Path to the private key file.")
	giconfig.Add(CAFile, "./cert/out/blackbox.crt", "Path to the certificate authority (CA).")

}
