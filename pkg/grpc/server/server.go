package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"

	c "github.com/b2wdigital/goignite/pkg/config"
	"github.com/b2wdigital/goignite/pkg/grpc/server/config"
	"github.com/b2wdigital/goignite/pkg/grpc/server/interceptor"
	"github.com/b2wdigital/goignite/pkg/log/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"
)

var	(
	instance *grpc.Server
)

func Start(ctx context.Context) *grpc.Server {

	log := logrus.FromContext(ctx)

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

	instance = s

	return instance
}


func Serve(ctx context.Context) {

	log := logrus.FromContext(ctx)

	service.RegisterChannelzServiceToServer(instance)

	// Register reflection service on gRPC server.
	reflection.Register(instance)

	port := c.Int(config.Port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	log.Infof("grpc server started on port %v", port)

	if err := instance.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}

}