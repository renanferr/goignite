package grpc

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strconv"

	"github.com/b2wdigital/goignite/pkg/health"
	"github.com/b2wdigital/goignite/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

func NewClient(ctx context.Context, options *Options) *grpc.ClientConn {

	var err error
	var conn *grpc.ClientConn
	var opts []grpc.DialOption

	l := log.FromContext(ctx)

	serverAddr := options.Host + ":" + strconv.Itoa(options.Port)

	if options.Tls {

		opts = addTlsOptions(ctx, options, opts)

	} else {

		opts = append(opts, grpc.WithInsecure())

	}

	if options.Gzip {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	}

	if options.HostOverwrite != "" {
		opts = append(opts, grpc.WithAuthority(options.HostOverwrite))
	}

	opts = append(opts, grpc.WithStreamInterceptor(DebugStreamClientInterceptor()))
	opts = append(opts, grpc.WithUnaryInterceptor(DebugUnaryClientInterceptor()))

	conn, err = grpc.Dial(serverAddr, opts...)

	if err != nil {
		l.Fatalf("fail to dial: %v", err)
		return nil
	}

	if options.Health.Enabled {
		configureHealthCheck(conn, options)
	}

	return conn
}

func configureHealthCheck(conn *grpc.ClientConn, o *Options) {
	cc := NewClientChecker(conn)
	hc := health.NewHealthChecker("grpc", o.Health.Description, cc, o.Health.Required)

	health.Add(hc)
}

func addTlsOptions(ctx context.Context, options *Options, opts []grpc.DialOption) []grpc.DialOption {

	l := log.FromContext(ctx)

	// Load the client certificates from disk
	cert, err := tls.LoadX509KeyPair(options.CertFile, options.KeyFile)
	if err != nil {
		l.Fatalf("could not load client key pair: %s", err)
	}
	var creds credentials.TransportCredentials
	if options.CAFile != "" {

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(options.CAFile)
		if err != nil {
			l.Fatalf("could not read ca certificate: %s", err)
		}

		// Append the certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			l.Fatalf("failed to append ca certs")
		}

		creds = credentials.NewTLS(&tls.Config{
			ServerName:         options.Host, // NOTE: this is required!
			Certificates:       []tls.Certificate{cert},
			RootCAs:            certPool,
			InsecureSkipVerify: options.InsecureSkipVerify,
		})

	} else {

		creds = credentials.NewTLS(&tls.Config{
			ServerName:         options.Host, // NOTE: this is required!
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: options.InsecureSkipVerify,
		})

	}
	opts = append(opts, grpc.WithTransportCredentials(creds))
	return opts
}
