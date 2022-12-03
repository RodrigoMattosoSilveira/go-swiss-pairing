package server

import (
	"context"
	"errors"
	"fmt"
	cm "github.com/RodrigoMattosoSilveira/go-swiss-pairing/grpc/server/interface/rpc/proto/swiss-pairing-apis/member/v1"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/usecase"
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
	cMember, cMemberError := c.useCase.Create(Member.First, Member.Last, Member.Email, Member.Password, Member.Cell)
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

func (c *MemberGrpcServer) Ping(_ context.Context, Member *cm.MemberPing) (*cm.MemberPong, error) {
	if Member.Ping != "ping" {
		message := fmt.Sprintf("Invalid ping paayload %s", Member.Ping)
		return nil, errors.New(message)
	}
	pong := &cm.MemberPong{
		Pong: "pong",
	}
	return pong, nil
}
