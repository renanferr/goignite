package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const (
	TlsEnabled = "grpc.server.tls.enabled"
	CertFile   = "grpc.server.tls.certfile"
	KeyFile    = "grpc.server.tls.keyfile"
	CaFile     = "grpc.server.tls.cafile"
)

func init() {
	log.Println("getting tls configurations")

	config.Add(TlsEnabled, false, "Use TLS - required for HTTP2.")
	config.Add(CertFile, "./cert/out/localhost.crt", "Path to the CRT/PEM file.")
	config.Add(KeyFile, "./cert/out/localhost.key", "Path to the private key file.")
	config.Add(CaFile, "./cert/out/blackbox.crt", "Path to the certificate authority (CA).")

}
