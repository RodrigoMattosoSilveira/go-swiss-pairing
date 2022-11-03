package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"

	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/constants"
	memberGrpc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var op string
	flag.StringVar(&op, "op", "ping", "the operation we want to execute")
	flag.Parse()

	type Member struct {
		First string
		Email string
	}

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

	// Perform the op
	switch op {
	case "ping":
		pong, err := client.Ping(ctx, &memberGrpc.MemberPing{Ping: "ping"})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(pong)
	case "create":
		newMember, err := client.Create(ctx, &memberGrpc.NewMember{First: "mario", Email: "mario@yahoo.com"})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created: id: %s name: %s email: %s\n", newMember.Id, newMember.First, newMember.Email)
	case "read":
		log.Println("Reading")
	case "readEmail":
		member, err := client.ReadEmail(ctx, &memberGrpc.MemberEmail{Email: "mario@yahoo.com"})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created: id: %s name: %s email: %s\n", member.Id, member.First, member.Email)
	case "readId":
		member, err := client.ReadId(ctx, &memberGrpc.MemberId{Id: "d67b65cf-5e50-4eae-94c8-3019c058c1d5"})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created: id: %s name: %s email: %s\n", member.Id, member.First, member.Email)
	default:
		log.Printf("Invalid operation: %s\n\n", op)
	}
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
