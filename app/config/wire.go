//go:build wireinject

package config

import (
	"github.com/google/wire"
)

func InitializeEvent() usecase.MemberUsecase {
	wire.Build(NewMemberRepository, NewMemberService, NewMemberUsecase, NewMemberGrpcServer)
	return NewMemberGrpcServer{}
}
