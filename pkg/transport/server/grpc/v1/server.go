package v1

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	"github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/log"
	rootgrpc "github.com/b2wdigital/goignite/pkg/transport/server/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

var (
	instance *grpc.Server
)

func Start(ctx context.Context) *grpc.Server {

	l := log.FromContext(ctx)

	gzip.SetLevel(9)

	var s *grpc.Server

	if config.Bool(rootgrpc.TlsEnabled) {

		// Load the certificates from disk
		certificate, err := tls.LoadX509KeyPair(config.String(rootgrpc.CertFile), config.String(rootgrpc.KeyFile))
		if err != nil {
			l.Fatalf("could not load server key pair: %s", err)
		}

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(config.String(rootgrpc.CaFile))
		if err != nil {
			l.Fatalf("could not read ca certificate: %s", err)
		}

		// Append the client certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			l.Fatalf("failed to append client certs")
		}

		// Create the TLS credentials
		creds := credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{certificate},
			ClientCAs:    certPool,
		})

		s = grpc.NewServer(
			grpc.Creds(creds),
			grpc.MaxConcurrentStreams(uint32(config.Int64(rootgrpc.MaxConcurrentStreams))),
			// grpc.InitialConnWindowSize(100),
			// grpc.InitialWindowSize(100),
			grpc.StreamInterceptor(DebugStreamServerInterceptor()),
			grpc.UnaryInterceptor(DebugUnaryServerInterceptor()),
		)

	} else {

		s = grpc.NewServer(
			grpc.MaxConcurrentStreams(uint32(config.Int64(rootgrpc.MaxConcurrentStreams))),
			// grpc.InitialConnWindowSize(100),
			// grpc.InitialWindowSize(100),
			grpc.StreamInterceptor(DebugStreamServerInterceptor()),
			grpc.UnaryInterceptor(DebugUnaryServerInterceptor()),
		)

	}

	instance = s

	return instance
}

func Serve(ctx context.Context) {

	l := log.FromContext(ctx)

	service.RegisterChannelzServiceToServer(instance)

	// Register reflection service on gRPC server.
	reflection.Register(instance)

	port := config.Int(rootgrpc.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		l.Fatalf("failed to listen: %v", err)
		return
	}

	l.Infof("grpc server started on port %v", port)

	if err := instance.Serve(lis); err != nil {
		l.Fatalf("failed to serve: %v", err)
		return
	}

}
