package service

import (
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/repository"
)

type ClubMemberService struct {
	repo repository.ClubMemberRepository
}

func NewClubMemberService(repo repository.ClubMemberRepository) *ClubMemberService {
	return &ClubMemberService{
		repo: repo,
	}
}

// DuplicatedEmail returns an error if there is a club member with this email, false otherwise
func (s *ClubMemberService) DuplicatedEmail(email string) bool {
	_, err := s.repo.ReadByEmail(email)
	if err != nil {
		return false
	}
	return true
}
