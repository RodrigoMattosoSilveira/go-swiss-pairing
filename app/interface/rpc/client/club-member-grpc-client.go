package client

import (
	"context"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
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

func (s *clubMemberGrpcService) ListClubMember(ctx context.Context, in *club_member_proto.ListClubMemberRequestType) (*club_member_proto.ListClubMemberResponseType, error) {
	cLubMembers, err := s.clubMemberUsecase.ListCLubMember()
	if err != nil {
		return nil, err
	}

	res := &club_member_proto.ListClubMemberResponseType{
		ClubMembers: toCLubMember(cLubMembers),
	}

	return res, nil
}

func (s *clubMemberGrpcService) RegisterClubMember(ctx context.Context, in *club_member_proto.RegisterClubMemberRequestType) (*club_member_proto.RegisterClubMemberResponseType, error) {
	if err := s.clubMemberUsecase.RegisterCLubMember(in.GetFirst(), in.GetEmail()); err != nil {
		return &club_member_proto.RegisterClubMemberResponseType{}, err
	}
	return &club_member_proto.RegisterClubMemberResponseType{}, nil
}

func toCLubMember(cLubMembers []*usecase.ClubMemberDAO) []*club_member_proto.ClubMember {
	res := make([]*club_member_proto.ClubMember, len(cLubMembers))
	for i, cLubMember := range cLubMembers {
		res[i] = &club_member_proto.ClubMember{
			Id:    cLubMember.Id,
			First: cLubMember.First,
			Email: cLubMember.Email,
		}
	}
	return res
}
