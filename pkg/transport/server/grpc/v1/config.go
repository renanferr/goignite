package grpc

import (
	"log"

	"github.com/b2wdigital/goignite/pkg/config"
)

const (
	Port                 = "transport.server.grpc.port"
	MaxConcurrentStreams = "transport.server.grpc.maxconcurrentstreams"
	TlsEnabled           = "transport.server.grpc.tls.enabled"
	CertFile             = "transport.server.grpc.tls.certfile"
	KeyFile              = "transport.server.grpc.tls.keyfile"
	CaFile               = "transport.server.grpc.tls.cafile"
)

func init() {

	log.Println("getting configurations for grpc server")

	config.Add(Port, 9090, "server grpc port")
	config.Add(MaxConcurrentStreams, 5000, "server grpc max concurrent streams")
	config.Add(TlsEnabled, false, "Use TLS - required for HTTP2.")
	config.Add(CertFile, "./cert/out/localhost.crt", "Path to the CRT/PEM file.")
	config.Add(KeyFile, "./cert/out/localhost.key", "Path to the private key file.")
	config.Add(CaFile, "./cert/out/blackbox.crt", "Path to the certificate authority (CA).")

}
