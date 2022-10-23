package main

import (
	"flag"
	"fmt"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/service"
	repo "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/persistence/memory"
	pb "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
	grpcServer "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/server"
	uc "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
	//"github.com/google/wire"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repository := repo.NewMemberRepository()
	svc := service.NewMemberService(repository)
	useCase := uc.NewMemberUsecase(repository, svc)
	server := grpcServer.NewMemberGrpcServer(useCase)

	s := grpc.NewServer()
	pb.RegisterMemberServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
