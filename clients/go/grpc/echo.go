package grpc

import (
	"github.com/jooita/GrpcRestExamples/echopb"
	"context"
	"crypto/x509"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func EchoInsecure(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := echopb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &echopb.StringMessage{Value: "Hello, World!"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("echo: %s", r.Value)
}

func Echo(address string, serverCrt []byte) {

	demoCertPool := x509.NewCertPool()
	ok := demoCertPool.AppendCertsFromPEM(serverCrt)
	if !ok {
		log.Fatal("bad certs")
	}

	var opts []grpc.DialOption
	creds := credentials.NewClientTLSFromCert(demoCertPool, address)
	opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := echopb.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Echo(ctx, &echopb.StringMessage{Value: "Hello, World!"})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("echo: %s", r.Value)
}
