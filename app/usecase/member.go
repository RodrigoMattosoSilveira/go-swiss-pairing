package usecase

import (
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/constants"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/model"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/repository"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/service"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"log"
	"regexp"
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
	Empty()
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
	//log.Printf("usecase/member/Create: called")

	// We must have a first name
	if first == "" {
		msg := "first not provided"
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}
	// We must have an email address
	if email == "" {
		msg := "email not provided"
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}
	// The email address must be valid
	//_, err := mail.ParseAddress(email)
	matchPattern := "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])"
	re, err := regexp.Compile(matchPattern)
	if err != nil {
		msg := "invalid regular expression to validate email address"
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}
	if !re.MatchString(email) {
		msg := "invalid email address"
		log.Printf(msg + ": " + email)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}
	// Ensure this is a unique email address
	if cm.service.DuplicatedEmail(email) {
		msg := "Duplicated email address: " + email
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}
	// Create a UUID for new member
	uid, err := uuid.NewRandom()
	if err != nil {
		log.Printf("unable to cread uuid for new member")
		return nil, status.Error(constants.GRPC_STATUS_ALREADY_EXISTS, err.Error())
	}
	// Create a Member Model
	MemberModel := model.Create(uid.String(), first, email)
	// Persist Member
	Member, err := cm.repo.Create(MemberModel)
	if err != nil {
		log.Printf("unable to create new member")
		return nil, status.Error(500, err.Error())
	}
	// Persisted new member
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

func (cm *MemberUsecase) Empty() {
	cm.repo.Empty()
}
