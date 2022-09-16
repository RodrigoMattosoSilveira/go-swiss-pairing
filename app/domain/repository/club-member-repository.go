package repository

import "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/model"

type ClubMemberRepository interface {
	// ReadAll find all club members
	Read() ([]*model.ClubMember, error)
	// ReadByEmail find a club member with the given email
	ReadByEmail(email string) (*model.ClubMember, error)
	// ReadById find a club member with the given email
	ReadById(id string) (*model.ClubMember, error)
	// Create club member with
	Create(*model.ClubMember) (*model.ClubMember, error)
}
