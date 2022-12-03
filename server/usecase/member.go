package usecase

import (
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/constants"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/domain/model"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/domain/repository"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/domain/service"
	"github.com/go-passwd/validator"
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
	Create(first string, last string, email string, password string, cell string) (*model.Member, error)
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
func (cm *MemberUsecase) Create(first string, last string, email string, password string, cell string) (*model.Member, error) {
	//log.Printf("usecase/member/Create: called")

	/*
	 *******************************************************************************************************************
	 * Create a new member UUID
	 *******************************************************************************************************************
	 */
	uid, err := uuid.NewRandom()
	if err != nil {
		log.Printf("unable to cread uuid for new member")
		return nil, status.Error(constants.GRPC_STATUS_ALREADY_EXISTS, err.Error())
	}

	/*
	 *******************************************************************************************************************
	 * Fist Name
	 *******************************************************************************************************************
	 */
	// We must have a first name
	if first == "" {
		msg := "first name not provided"
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}

	/*
	 *******************************************************************************************************************
	 * Last Name
	 *******************************************************************************************************************
	 */
	// We must have a last name
	if last == "" {
		msg := "last name not provided"
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}

	/*
	 *******************************************************************************************************************
	 * Email Address
	 *******************************************************************************************************************
	 */
	// We must have an email address
	if email == "" {
		msg := "email not provided"
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}
	// The email address must be valid
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
	// The email address must be unique
	if cm.service.DuplicatedEmail(email) {
		msg := "Duplicated email address: " + email
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}

	/*
	 *******************************************************************************************************************
	 * Password
	 *******************************************************************************************************************
	 */
	// We must have a password
	if password == "" {
		msg := "password not provided"
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}
	// Validate Password, it must:
	// contain at least 8 characters and at most 20 characters.
	// contain at least one digit.
	// contain at least one upper case alphabet.
	// contain at least one lower case alphabet.
	// contains at least one special character which includes !@#$%&*()-+=^.
	// does not contain any white space.
	//ok, err := validatePassword(password)
	if !ValidatePassword(password) {
		msg := "invalid password"
		log.Printf(msg + ": " + password)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}

	/*
	 *******************************************************************************************************************
	 * Cell phone number
	 *******************************************************************************************************************
	 */
	// We must have a Cell phone number
	if cell == "" {
		msg := "cell phone number not provided"
		log.Printf(msg)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}
	if !ValidateUsaCellPhoneNumber(cell) {
		msg := "invalid cell phone number"
		log.Printf(msg + ": " + cell)
		return nil, status.Error(constants.GRPC_STATUS_INVALID_ARGUMENT, msg)
	}

	// Create a Member Model
	MemberModel := model.Create(
		uid.String(),
		first,
		last,
		email,
		password,
		cell,
		1200,
		true,
		"picture_placeholder.jpeg")

	// Persist new Member
	log.Printf("creating new member")
	Member, err := cm.repo.Create(MemberModel)
	if err != nil {
		log.Printf("unable to create new member")
		return nil, status.Error(500, err.Error())
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
		log.Printf(err.Error())
		return nil, status.Error(constants.GRPC_STATUS_NOT_FOUND, err.Error())
	}
	return Member, nil
}

func (cm *MemberUsecase) ReadById(id string) (*model.Member, error) {
	Member, err := cm.repo.ReadById(id)
	if err != nil {
		log.Printf(err.Error())
		return nil, status.Error(constants.GRPC_STATUS_NOT_FOUND, err.Error())
	}
	return Member, nil
}

func (cm *MemberUsecase) Empty() {
	cm.repo.Empty()
}

// TODO perhaps this should be in an utilities package
func ValidateUsaCellPhoneNumber(cellPhoneNumber string) bool {
	reCell, errCell := regexp.Compile(`^\(?(\d{3})\)?[-\. ]?(\d{3})[-\. ]?(\d{4})( x\d{4})?$`)
	if errCell != nil {
		msg := "invalid cell phone number regex"
		log.Printf(msg)
		return false
	}
	if !reCell.MatchString(cellPhoneNumber) {
		msg := "invalid cell phone number"
		log.Printf(msg + ": " + cellPhoneNumber)
		return false
	}
	return true

}

// TODO perhaps this should be in an utilities package
func ValidatePassword(password string) bool {
	// contain at least 8 characters and at most 20 characters.
	passwordValidator := validator.New(validator.MinLength(8, nil), validator.MaxLength(20, nil))
	err := passwordValidator.Validate(password)
	if err != nil {
		log.Printf("Invalid password length")
		return false
	}

	// contain at least one lower case character.
	passwordValidator = validator.New(validator.ContainsAtLeast(`abcdefghijklmnopqrstuvwxyz`, 1, nil))
	err = passwordValidator.Validate(password)
	if err != nil {
		log.Printf("Missing at least one lower case character")
		return false
	}

	// contain at least one upper case character.
	passwordValidator = validator.New(validator.ContainsAtLeast(`ABCDEFGHIJKLMNOPQRSTUVWXYZ`, 1, nil))
	err = passwordValidator.Validate(password)
	if err != nil {
		log.Printf("Missing at least one upper case character")
		return false
	}

	// contain at least one digit.
	passwordValidator = validator.New(validator.ContainsAtLeast(`1234567890`, 1, nil))
	err = passwordValidator.Validate(password)
	if err != nil {
		log.Printf("Missing at least one digit")
		return false
	}

	// contain at least one special character.
	passwordValidator = validator.New(validator.ContainsAtLeast(`?=.*[@#$%^&-+=()]`, 1, nil))
	err = passwordValidator.Validate(password)
	if err != nil {
		log.Printf("Missing at least one special character")
		return false
	}

	// does not contain any white space.
	passwordValidator = validator.New(validator.ContainsAtLeast(` `, 1, nil))
	err = passwordValidator.Validate(password)
	if err == nil {
		log.Printf("Have at least one invalid space")
		return false
	}

	return true
}
