//go:build wireinject

package config

import (
	"github.com/google/wire"
)

func InitializeEvent() usecase.ClubMemberUsecase {
	wire.Build(NewClubMemberRepository, NewClubMemberService, NewClubMemberUsecase, NewClubMemberGrpcServer)
	return NewClubMemberGrpcServer{}
}
