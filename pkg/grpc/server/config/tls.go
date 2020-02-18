package config

import (
	"log"

	"github.com/jpfaria/goignite/pkg/config"
)

const TlsEnabled = "server.grpc.tls.enabled"
const CertFile = "server.grpc.tls.certfile"
const KeyFile = "server.grpc.tls.keyfile"
const CaFile = "server.grpc.tls.cafile"

func init() {
	log.Println("getting tls configurations")

	config.SetBool(TlsEnabled, false, "Use TLS - required for HTTP2.")
	config.SetString(CertFile, "./cert/out/localhost.crt", "Path to the CRT/PEM file.")
	config.SetString(KeyFile, "./cert/out/localhost.key", "Path to the private key file.")
	config.SetString(CaFile, "./cert/out/blackbox.crt", "Path to the certificate authority (CA).")

}
