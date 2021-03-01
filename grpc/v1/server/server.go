package gigrpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
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

	logger := gilog.FromContext(ctx)

	gzip.SetLevel(9)

	var s *grpc.Server

	if giconfig.Bool(TlsEnabled) {

		// Load the certificates from disk
		certificate, err := tls.LoadX509KeyPair(giconfig.String(CertFile), giconfig.String(KeyFile))
		if err != nil {
			logger.Fatalf("could not load server key pair: %s", err)
		}

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(giconfig.String(CAFile))
		if err != nil {
			logger.Fatalf("could not read ca certificate: %s", err)
		}

		// Append the client certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			logger.Fatalf("failed to append client certs")
		}

		// Create the TLS credentials
		creds := credentials.NewTLS(&tls.Config{
			ClientAuth:   tls.RequireAndVerifyClientCert,
			Certificates: []tls.Certificate{certificate},
			ClientCAs:    certPool,
		})

		s = grpc.NewServer(
			grpc.Creds(creds),
			grpc.MaxConcurrentStreams(uint32(giconfig.Int64(MaxConcurrentStreams))),
			// grpc.InitialConnWindowSize(100),
			// grpc.InitialWindowSize(100),
			grpc.StreamInterceptor(DebugStreamServerInterceptor()),
			grpc.UnaryInterceptor(DebugUnaryServerInterceptor()),
		)

	} else {

		s = grpc.NewServer(
			grpc.MaxConcurrentStreams(uint32(giconfig.Int64(MaxConcurrentStreams))),
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

	logger := gilog.FromContext(ctx)

	service.RegisterChannelzServiceToServer(instance)

	// Register reflection service on gRPC server.
	reflection.Register(instance)

	port := giconfig.Int(Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
		return
	}

	logger.Infof("grpc server started on port %v", port)

	if err := instance.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
		return
	}

}
