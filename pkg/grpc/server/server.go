package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	c "github.com/jpfaria/goignite/pkg/config"
	"github.com/jpfaria/goignite/pkg/grpc/server/config"
	"github.com/jpfaria/goignite/pkg/grpc/server/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpc *grpc.Server
}

func (s *Server) Start() {

	log.Println("grpc server starting")

	service.RegisterChannelzServiceToServer(s.grpc)

	// Register reflection service on gRPC server.
	reflection.Register(s.grpc)

	port := c.Int(config.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	if err := s.grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}

}

func (s *Server) GetInstance() *grpc.Server {
	return s.grpc
}

func New() *Server {

	c.Parse()

	gzip.SetLevel(9)

	var s *grpc.Server

	if c.Bool(config.TlsEnabled) {

		// Load the certificates from disk
		certificate, err := tls.LoadX509KeyPair(c.String(config.CertFile), c.String(config.KeyFile))
		if err != nil {
			log.Fatalf("could not load server key pair: %s", err)
		}

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(c.String(config.CaFile))
		if err != nil {
			log.Fatalf("could not read ca certificate: %s", err)
		}

		// Append the client certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			log.Fatal("failed to append client certs")
		}

		// Create the TLS credentials
		creds := credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{certificate},
			ClientCAs:    certPool,
		})

		s = grpc.NewServer(
			grpc.Creds(creds),
			grpc.MaxConcurrentStreams(uint32(c.Int64(config.MaxConcurrentStreams))),
			// grpc.InitialConnWindowSize(100),
			// grpc.InitialWindowSize(100),
			grpc.StreamInterceptor(interceptor.DebugStreamServerInterceptor()),
			grpc.UnaryInterceptor(interceptor.DebugUnaryServerInterceptor()),
		)

	} else {

		s = grpc.NewServer(
			grpc.MaxConcurrentStreams(uint32(c.Int64(config.MaxConcurrentStreams))),
			// grpc.InitialConnWindowSize(100),
			// grpc.InitialWindowSize(100),
			grpc.StreamInterceptor(interceptor.DebugStreamServerInterceptor()),
			grpc.UnaryInterceptor(interceptor.DebugUnaryServerInterceptor()),
		)

	}

	return &Server{grpc: s}
}
