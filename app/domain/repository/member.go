package repository

import "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/model"

type MemberRepository interface {
	// Read find all members
	Read() ([]*model.Member, error)
	// ReadByEmail find a member with the given email
	ReadByEmail(email string) (*model.Member, error)
	// ReadById find a member with the given email
	ReadById(id string) (*model.Member, error)
	// Create member with
	Create(*model.Member) (*model.Member, error)
}
