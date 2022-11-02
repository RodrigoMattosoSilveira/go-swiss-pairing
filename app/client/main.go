package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/constants"
	memberGrpc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	ctx := context.Background()
	// Load our TLS certificate and use grpc/credentials to create new transport credentials
	creds := credentials.NewTLS(loadTLSCfg())
	// Create a new connection using the transport credentials
	conn, err := grpc.DialContext(ctx, "localhost:9990", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// A new GRPC client to use
	client := memberGrpc.NewMemberServiceClient(conn)

	pong, err := client.Ping(ctx, &memberGrpc.MemberPing{Ping: "ping"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pong)
}

// loadTLSCfg will load a certificate and create a tls config
func loadTLSCfg() *tls.Config {
	b, _ := ioutil.ReadFile(constants.ServerCertPem)
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		log.Fatal("credentials: failed to append certificates")
	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}
	return config
}