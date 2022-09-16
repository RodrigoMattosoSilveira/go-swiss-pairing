package usecase

import (
	"fmt"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/model"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/repository"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/service"
	"github.com/google/uuid"
)

// ClubMemberDAO an internal structure to construct a struct with attributes only. Note that this is similar, but it is
// not model.ClubMember; we want to ensure that the MODEL does not do know anything about its upper layers.
// TODO I wish there is a more elegant way to do this
//type ClubMemberDAO struct {
//	Id    string
//	First string
//	Email string
//}

type IClubMemberUsecase interface {
	Create(first string, email string) (*model.ClubMember, error)
	Read() ([]*model.ClubMember, error)
	ReadById(id string) (*model.ClubMember, error)
	ReadByEmail(email string) (*model.ClubMember, error)
}

type ClubMemberUsecase struct {
	repo    repository.ClubMemberRepository
	service *service.ClubMemberService
}

func NewClubMemberUsecase(repo repository.ClubMemberRepository, service *service.ClubMemberService) *ClubMemberUsecase {
	return &ClubMemberUsecase{
		repo:    repo,
		service: service,
	}
}

// Create Returns *ClubMemberDAO or error; when an error occurs, *ClubMemberDAO is nil
//
func (cm *ClubMemberUsecase) Create(first, email string) (*model.ClubMember, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	if cm.service.DuplicatedEmail(email) {
		return nil, fmt.Errorf("club-member-usecase/Create: Club Member with Email: %s already exists", email)
	}
	clubMemberModel := model.NewClubMember(uid.String(), first, email)
	clubMember, clubMemberErr := cm.repo.Create(clubMemberModel)
	if clubMemberErr != nil {
		return nil, clubMemberErr
	}
	return clubMember, nil
}

func (cm *ClubMemberUsecase) Read() ([]*model.ClubMember, error) {
	clubMembers, err := cm.repo.Read()
	if err != nil {
		return nil, err
	}
	return clubMembers, nil
}

func (cm *ClubMemberUsecase) ReadByEmail(email string) (*model.ClubMember, error) {
	clubMember, err := cm.repo.ReadByEmail(email)
	if err != nil {
		return nil, err
	}
	return clubMember, nil
}

func (cm *ClubMemberUsecase) ReadById(id string) (*model.ClubMember, error) {
	clubMember, err := cm.repo.ReadById(id)
	if err != nil {
		return nil, err
	}
	return clubMember, nil
}
