package client

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strconv"

	"github.com/jpfaria/goignite/pkg/grpc/client/interceptor"
	"github.com/jpfaria/goignite/pkg/grpc/client/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/encoding/gzip"
)

func NewConnection(options *model.Options) *grpc.ClientConn {

	var err error
	var conn *grpc.ClientConn
	var opts []grpc.DialOption

	serverAddr := options.Host + ":" + strconv.Itoa(options.Port)

	if options.Tls {

		opts = addTlsOptions(options, opts)

	} else {

		opts = append(opts, grpc.WithInsecure())

	}

	if options.Gzip {
		opts = append(opts, grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)))
	}

	if options.HostOverwrite != "" {
		opts = append(opts, grpc.WithAuthority(options.HostOverwrite))
	}

	opts = append(opts, grpc.WithStreamInterceptor(interceptor.DebugStreamClientInterceptor()))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor.DebugUnaryClientInterceptor()))

	conn, err = grpc.Dial(serverAddr, opts...)

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return nil
	}

	return conn

}

func addTlsOptions(options *model.Options, opts []grpc.DialOption) []grpc.DialOption {
	// Load the client certificates from disk
	cert, err := tls.LoadX509KeyPair(options.CertFile, options.KeyFile)
	if err != nil {
		log.Fatalf("could not load client key pair: %s", err)
	}
	var creds credentials.TransportCredentials
	if options.CAFile != "" {

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(options.CAFile)
		if err != nil {
			log.Fatalf("could not read ca certificate: %s", err)
		}

		// Append the certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			log.Fatal("failed to append ca certs")
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
