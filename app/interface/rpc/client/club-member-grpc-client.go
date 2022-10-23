package client

import (
	"context"
	pb "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/usecase"
)

type clubMemberGrpcService struct {
	clubMemberUsecase usecase.ClubMemberUsecase
}

func NewClubMemberGrpcService(clubMemberUsecase usecase.ClubMemberUsecase) *clubMemberGrpcService {
	return &clubMemberGrpcService{
		clubMemberUsecase: clubMemberUsecase,
	}
}

func (s *clubMemberGrpcService) Create(ctx context.Context, in *pb.NewClubMember) (*pb.NewClubMember, error) {
	if err, _ := s.clubMemberUsecase.Create(in.GetFirst(), in.GetEmail()); err != nil {
		return nil, nil
	}
	return &pb.NewClubMember{}, nil
}
