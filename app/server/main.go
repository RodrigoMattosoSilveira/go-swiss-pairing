package main

import (
	"flag"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/constants"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/service"
	repo "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/persistence/memory"
	memberGrpc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
	pb "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/server"
	uc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/usecase"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"net/http"
	"time"
	//"github.com/google/wire"
)

var (
	tcpPort  = flag.Int("tcpPort", 9990, "The server tcpPort")
	httpPort = flag.Int("httpPort", 8080, "The client httpPort")
)

func main() {

	// We Generate a TLS grpc API
	grpcServer, err := GenerateTLSApi(constants.ServerCertPem, constants.ServerKeyPem)
	if err != nil {
		log.Fatal(err)
	}
	// Start listening on a TCP Port
	tcpListener, err := net.Listen("tcp", "127.0.0.1:9990")
	if err != nil {
		log.Fatal(err)
	}
	// We need to tell the code WHAT TO do on each request, ie. The business logic.
	// In GRPC cases, the Server is acutally just an Interface
	// So we need a struct which fulfills the server interface
	// see server.go
	repository := repo.NewMemberRepository()
	svc := service.NewMemberService(repository)
	useCase := uc.NewMemberUsecase(repository, svc)
	memberServiceServer := pb.NewMemberGrpcServer(useCase)
	// Register the API server as a PingPong Server
	// The register function is a generated piece by protoc.
	memberGrpc.RegisterMemberServiceServer(grpcServer, memberServiceServer)

	// Start serving in a goroutine to not block
	go func() {
		log.Fatal(grpcServer.Serve(tcpListener))
	}()
	// Wrap the GRPC Server in grpc-web and also host the UI
	grpcWebServer := grpcweb.WrapServer(grpcServer)
	// Lets put the wrapped grpc server in our multiplexer struct so
	// it can reach the grpc server in its handler
	multiplex := grpcMultiplexer{
		grpcWebServer,
	}

	// We need a http router
	r := http.NewServeMux()
	// Load the static webpage with a http fileserver
	webapp := http.FileServer(http.Dir("ui/pingpongapp/build"))
	// Host the Web Application at /, and wrap it in the GRPC Multiplexer
	// This allows grpc requests to transfer over HTTP1. then be
	// routed by the multiplexer
	r.Handle("/", multiplex.Handler(webapp))
	// Create a HTTP server and bind the router to it, and set wanted address
	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	// Serve the webapp over TLS
	log.Printf("Starting server on localhost:8080")
	log.Fatal(srv.ListenAndServeTLS(constants.ServerCertPem, constants.ServerKeyPem))

}

// GenerateTLSApi will load TLS certificates and key and create a grpc server with those.
func GenerateTLSApi(pemPath, keyPath string) (*grpc.Server, error) {
	cred, err := credentials.NewServerTLSFromFile(pemPath, keyPath)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(
		grpc.Creds(cred),
	)
	return s, nil
}

type grpcMultiplexer struct {
	*grpcweb.WrappedGrpcServer
}

// Handler is used to route requests to either grpc or to regular http
func (m *grpcMultiplexer) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
