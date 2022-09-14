package memory

import (
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/model"
	"sync"
)

// It
type clubMemberMemory struct {
	Id    string
	First string
	Email string
}

type clubMemberRepository struct {
	mu          *sync.Mutex
	clubMembers map[string]*clubMemberMemory
}

func NewClubMemberRepository() *clubMemberRepository {
	return &clubMemberRepository{
		mu:          &sync.Mutex{},
		clubMembers: map[string]*clubMemberMemory{},
	}
}

func (r *clubMemberRepository) ReadAll() ([]*model.ClubMember, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]*model.ClubMember, len(r.clubMembers))
	i := 0
	for _, clubMember := range r.clubMembers {
		users[i] = model.NewClubMember(clubMember.Id, clubMember.First, clubMember.Email)
		i++
	}
	return users, nil
}

func (r *clubMemberRepository) ReadById(id string) (*model.ClubMember, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, clubMember := range r.clubMembers {
		if clubMember.Id == id {
			return model.NewClubMember(clubMember.Id, clubMember.First, clubMember.Email), nil
		}
	}
	return nil, nil
}

func (r *clubMemberRepository) ReadByEmail(email string) (*model.ClubMember, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, clubMember := range r.clubMembers {
		if clubMember.Email == email {
			return model.NewClubMember(clubMember.Id, clubMember.First, clubMember.Email), nil
		}
	}
	return nil, nil
}
func (r *clubMemberRepository) Create(clubMember *model.ClubMember) (*model.ClubMember, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.clubMembers[clubMember.Id()] = &clubMemberMemory{
		Id:    clubMember.Id(),
		First: clubMember.First(),
		Email: clubMember.Email(),
	}
	return clubMember, nil
}
