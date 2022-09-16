package server

import (
	"context"
	cm "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/usecase"
	"sync"
)

type clubMemberGrpcServer struct {
	mu      sync.Mutex // protects routeNotes
	useCase *usecase.ClubMemberUsecase
}

func NewClubMemberGrpcServer(useCase *usecase.ClubMemberUsecase) *clubMemberGrpcServer {
	return &clubMemberGrpcServer{
		useCase: useCase,
	}
}

func (c *clubMemberGrpcServer) Create(_ context.Context, clubMember *cm.NewClubMember) (*cm.ClubMember, error) {
	cMember, cMemberError := c.useCase.Create(clubMember.First, clubMember.Email)
	if cMemberError != nil {
		return nil, cMemberError
	}
	createClubMember := &cm.ClubMember{
		Id:    cMember.Id(),
		First: cMember.First(),
		Email: cMember.Email(),
	}
	return createClubMember, nil
}

func (c *clubMemberGrpcServer) Read(_ *cm.ClubMemberEmpty, stream cm.ClubMemberService_ReadServer) error {
	cMembers, cMemberError := c.useCase.Read()
	if cMemberError != nil {
		return cMemberError
	}
	for _, cMember := range cMembers {
		streamClubMember := &cm.ClubMember{
			Id:    cMember.Id(),
			First: cMember.First(),
			Email: cMember.Email(),
		}
		if err := stream.Send(streamClubMember); err != nil {
			return err
		}
	}
	return nil
}

func (c *clubMemberGrpcServer) ReadEmail(_ context.Context, clubMember *cm.ClubMemberEmail) (*cm.ClubMember, error) {
	cMember, cMemberError := c.useCase.ReadByEmail(clubMember.Email)
	if cMemberError != nil {
		return nil, cMemberError
	}
	var readClubMemberEmail = &cm.ClubMember{
		Id:    cMember.Id(),
		First: cMember.First(),
		Email: cMember.Email(),
	}
	return readClubMemberEmail, nil
}

func (c *clubMemberGrpcServer) ReadId(_ context.Context, clubMember *cm.ClubMemberId) (*cm.ClubMember, error) {
	cMember, cMemberError := c.useCase.ReadById(clubMember.Id)
	if cMemberError != nil {
		return nil, cMemberError
	}
	var readClubMemberEmail = &cm.ClubMember{
		Id:    cMember.Id(),
		First: cMember.First(),
		Email: cMember.Email(),
	}
	return readClubMemberEmail, nil
}
