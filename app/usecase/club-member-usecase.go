package usecase

import (
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
	ReadAll() ([]*model.ClubMember, error)
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
	if err := cm.service.Duplicated(email); err != nil {
		return nil, err
	}
	clubMemberModel := model.NewClubMember(uid.String(), first, email)
	clubMember, clubMemberErr := cm.repo.Create(clubMemberModel)
	if clubMember != nil {
		return clubMember, clubMemberErr
	}
	return nil, clubMemberErr
}

func (cm *ClubMemberUsecase) ReadAll() ([]*model.ClubMember, error) {
	clubMembers, err := cm.repo.ReadAll()
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
	clubMember, err := cm.repo.ReadByEmail(id)
	if err != nil {
		return nil, err
	}
	return clubMember, nil
}

//
//func toCLubMember(clubMembers []*model.ClubMember) []*model.ClubMember {
//	res := make([]*model.ClubMember, len(clubMembers))
//	for i, clubMember := range clubMembers {
//		res[i] = fromModelToDao(clubMember)
//	}
//	return res
//}
//
//func fromModelToDao(cm *model.ClubMember) *model.ClubMember {
//	return &model.ClubMember{
//		id:    cm.Id(),
//		first: cm.First(),
//		email: cm.Email(),
//	}
//}
