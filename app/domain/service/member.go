package service

import (
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/repository"
)

type MemberService struct {
	repo repository.MemberRepository
}

func NewMemberService(repo repository.MemberRepository) *MemberService {
	return &MemberService{
		repo: repo,
	}
}

// DuplicatedEmail returns an error if there is a club member with this email, false otherwise
func (s *MemberService) DuplicatedEmail(email string) bool {
	_, err := s.repo.ReadByEmail(email)
	if err != nil {
		return false
	}
	return true
}
