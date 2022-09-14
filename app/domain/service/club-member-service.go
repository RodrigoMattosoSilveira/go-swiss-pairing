package service

import (
	"fmt"
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
func (s *ClubMemberService) Duplicated(email string) error {
	clubMember, err := s.repo.ReadByEmail(email)
	if clubMember != nil {
		return fmt.Errorf("%s already exists", email)
	}
	if err != nil {
		return err
	}
	return nil
}
