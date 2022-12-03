package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"

	memberGrpc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/grpc/server/interface/rpc/proto/swiss-pairing-apis/member/v1"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/*
 Usage: ./server/client/main --op ping
         ./client/main --op create --first <<first>>  --last <<last>> --email <<email>> --password <<password>> --cell <<cell>>
		 ./client/main --op read
		 ./client/main --op readEmail --email <<email>>
		 ./client/main --op readId --id <<id>>
		 ./client/main --op seed --seedData <<folder name>>
*/

type Member struct {
	Id       string
	First    string
	Last     string
	Email    string
	Password string
	Cell     string
	Rating   int32
	IsActive bool
	ImageUrl string
}

func main() {
	var op, id, first, last, email, password, cell, seedData string
	flag.StringVar(&op, "op", "ping", "the operation we want to execute")
	flag.StringVar(&id, "id", "", "the member id")
	flag.StringVar(&first, "first", "", "the member first name")
	flag.StringVar(&last, "last", "", "the member last name")
	flag.StringVar(&email, "email", "", "the member email address")
	flag.StringVar(&password, "password", "", "the member password name")
	flag.StringVar(&cell, "cell", "", "the member cell phone number")
	flag.StringVar(&seedData, "seedData", "", "the seed data filename")

	flag.Parse()

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
		newMember, err := clientMemberCreate(client, ctx, first, last, email, password, cell)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created: id: %s name: %s email: %s\n", newMember.Id, newMember.First, newMember.Email)
	case "read":
		stream, err := client.Read(context.Background(), &memberGrpc.MemberEmpty{})
		if err != nil {
			log.Fatal(err)
		}
		// see https://linuxhint.com/golang-make-function/ for a brief discussion on how to use make to create a channel
		// for to connect concurrent goroutines
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
	case "seed":
		log.Printf("Seeding the in-memory database")
		if seedData == "NONE" {
			log.Fatal("Seed data file not provided")
		}
		if _, err := os.Stat(seedData); err == nil {
			log.Printf("Found seedData: %s \n", seedData)
		} else {
			log.Fatalf("Did not find seedData: %s", seedData)
		}
		// Read the seedData file
		fileContent, err := os.Open(seedData)
		if err != nil {
			log.Fatal("Error opening seedData file: ", err)
		}
		defer fileContent.Close()
		seedDataContent, err := ioutil.ReadFile(seedData)
		if err != nil {
			log.Fatal("Error reading seedData file: ", err)
		}

		// Unmarshall the data into `payload`
		var payload []Member
		err = json.Unmarshal(seedDataContent, &payload)
		if err != nil {
			log.Fatal("Error Unmarshalling seedData: ", err)
		}

		// Seed the memory database
		for _, member := range payload {
			// the value of v is assigned to a new local variable v
			member := member
			log.Printf("Seeding member: is %s\n", member.Email)
			newMember, error := clientMemberCreate(client, ctx, member.First, member.Last, member.Email, member.Password, member.Cell)
			if error != nil {
				log.Printf("Error seeding member: %s", member.Email)
				log.Fatalf(error.Error())
			}
			log.Printf("Success seeding member id: is %s\n", newMember.Email)
		}
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

func clientMemberCreate(client memberGrpc.MemberServiceClient, ctx context.Context, first string, last string, email string, password string, cell string) (*memberGrpc.Member, error) {
	//log.Printf("client/main/clientMemberCreate: called")
	member, error := client.Create(ctx, &memberGrpc.NewMember{First: first, Last: last, Email: email, Password: password, Cell: cell})
	if error != nil {
		return nil, error
	}
	log.Printf("Sucess seeding member: %s", email)
	return member, nil
}
