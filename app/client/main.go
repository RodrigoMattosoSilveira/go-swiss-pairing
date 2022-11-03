package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io"
	"io/ioutil"
	"log"

	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/constants"
	memberGrpc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var op, id, name, email string
	flag.StringVar(&op, "op", "ping", "the operation we want to execute")
	flag.StringVar(&id, "id", "NONE", "the member id")
	flag.StringVar(&name, "name", "NONE", "the member name")
	flag.StringVar(&email, "email", "NONE", "the member email")
	flag.Parse()
	//
	//type Member struct {
	//	First string
	//	Email string
	//}

	ctx := context.Background()
	// Load our TLS certificate and use grpc/credentials to create new transport credentials
	creds := credentials.NewTLS(loadTLSCfg())
	// Create a new connection using the transport credentials
	conn, err := grpc.DialContext(ctx, "localhost:9990", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)
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
		if name == "NONE" {
			log.Fatal("Name not provided")
		}
		if email == "NONE" {
			log.Fatal("Email not provided")
		}
		newMember, err := client.Create(ctx, &memberGrpc.NewMember{First: name, Email: email})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created: id: %s name: %s email: %s\n", newMember.Id, newMember.First, newMember.Email)
	case "read":
		stream, err := client.Read(context.Background(), &memberGrpc.MemberEmpty{})
		if err != nil {
			log.Fatal(err)
		}
		done := make(chan bool)
		go func() {
			for {
				member, err := stream.Recv()
				if err == io.EOF {
					done <- true //means stream is finished
					return
				}
				if err != nil {
					log.Fatalf("cannot receive %v", err)
				}
				log.Printf("Read member / id: %s, name: %s, email: %s\n", member.Id, member.First, member.Email)
			}
		}()
		<-done //we will wait until all response is received
	case "readEmail":
		if email == "NONE" {
			log.Fatal("Email not provided")
		}
		member, err := client.ReadEmail(ctx, &memberGrpc.MemberEmail{Email: email})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created: id: %s name: %s email: %s\n", member.Id, member.First, member.Email)
	case "readId":
		if id == "NONE" {
			log.Fatal("Id not provided")
		}
		member, err := client.ReadId(ctx, &memberGrpc.MemberId{Id: id})
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
