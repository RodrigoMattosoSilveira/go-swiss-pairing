package usecase

import (
	"fmt"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/model"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/repository"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/service"
	"github.com/google/uuid"
)

// MemberDAO an internal structure to construct a struct with attributes only. Note that this is similar, but it is
// not model.Member; we want to ensure that the MODEL does not do know anything about its upper layers.
// TODO I wish there is a more elegant way to do this
//type MemberDAO struct {
//	Id    string
//	First string
//	Email string
//}

type IMemberUsecase interface {
	Create(first string, email string) (*model.Member, error)
	Read() ([]*model.Member, error)
	ReadById(id string) (*model.Member, error)
	ReadByEmail(email string) (*model.Member, error)
}

type MemberUsecase struct {
	repo    repository.MemberRepository
	service *service.MemberService
}

func NewMemberUsecase(repo repository.MemberRepository, service *service.MemberService) *MemberUsecase {
	return &MemberUsecase{
		repo:    repo,
		service: service,
	}
}

// Create Returns *MemberDAO or error; when an error occurs, *MemberDAO is nil
//
func (cm *MemberUsecase) Create(first, email string) (*model.Member, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	if cm.service.DuplicatedEmail(email) {
		return nil, fmt.Errorf("member-usecase/Create: Member with Email: %s already exists", email)
	}
	MemberModel := model.Create(uid.String(), first, email)
	Member, MemberErr := cm.repo.Create(MemberModel)
	if MemberErr != nil {
		return nil, MemberErr
	}
	return Member, nil
}

func (cm *MemberUsecase) Read() ([]*model.Member, error) {
	Members, err := cm.repo.Read()
	if err != nil {
		return nil, err
	}
	return Members, nil
}

func (cm *MemberUsecase) ReadByEmail(email string) (*model.Member, error) {
	Member, err := cm.repo.ReadByEmail(email)
	if err != nil {
		return nil, err
	}
	return Member, nil
}

func (cm *MemberUsecase) ReadById(id string) (*model.Member, error) {
	Member, err := cm.repo.ReadById(id)
	if err != nil {
		return nil, err
	}
	return Member, nil
}
