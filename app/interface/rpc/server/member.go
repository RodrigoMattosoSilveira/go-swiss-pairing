package server

import (
	"context"
	cm "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/interface/rpc/proto"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/usecase"
	"sync"
)

type MemberGrpcServer struct {
	mu      sync.Mutex // protects routeNotes
	useCase *usecase.MemberUsecase
}

func NewMemberGrpcServer(useCase *usecase.MemberUsecase) *MemberGrpcServer {
	return &MemberGrpcServer{
		useCase: useCase,
	}
}

func (c *MemberGrpcServer) Create(_ context.Context, Member *cm.NewMember) (*cm.Member, error) {
	cMember, cMemberError := c.useCase.Create(Member.First, Member.Email)
	if cMemberError != nil {
		return nil, cMemberError
	}
	createMember := &cm.Member{
		Id:    cMember.Id(),
		First: cMember.First(),
		Email: cMember.Email(),
	}
	return createMember, nil
}

func (c *MemberGrpcServer) Read(_ *cm.MemberEmpty, stream cm.MemberService_ReadServer) error {
	cMembers, cMemberError := c.useCase.Read()
	if cMemberError != nil {
		return cMemberError
	}
	for _, cMember := range cMembers {
		streamMember := &cm.Member{
			Id:    cMember.Id(),
			First: cMember.First(),
			Email: cMember.Email(),
		}
		if err := stream.Send(streamMember); err != nil {
			return err
		}
	}
	return nil
}

func (c *MemberGrpcServer) ReadEmail(_ context.Context, Member *cm.MemberEmail) (*cm.Member, error) {
	cMember, cMemberError := c.useCase.ReadByEmail(Member.Email)
	if cMemberError != nil {
		return nil, cMemberError
	}
	var readMemberEmail = &cm.Member{
		Id:    cMember.Id(),
		First: cMember.First(),
		Email: cMember.Email(),
	}
	return readMemberEmail, nil
}

func (c *MemberGrpcServer) ReadId(_ context.Context, Member *cm.MemberId) (*cm.Member, error) {
	cMember, cMemberError := c.useCase.ReadById(Member.Id)
	if cMemberError != nil {
		return nil, cMemberError
	}
	var readMemberEmail = &cm.Member{
		Id:    cMember.Id(),
		First: cMember.First(),
		Email: cMember.Email(),
	}
	return readMemberEmail, nil
}
