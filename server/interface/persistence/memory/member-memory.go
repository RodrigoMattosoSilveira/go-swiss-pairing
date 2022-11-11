package memory

import (
	"errors"
	"fmt"
	"github.com/RodrigoMattosoSilveira/go-swiss-pairing/server/domain/model"
	"sync"
)

// It
type MemberMemory struct {
	Id    string
	First string
	Email string
}

type MemberRepository struct {
	mu      *sync.Mutex
	Members map[string]*MemberMemory
}

func NewMemberRepository() *MemberRepository {
	return &MemberRepository{
		mu:      &sync.Mutex{},
		Members: map[string]*MemberMemory{},
	}
}

func (r *MemberRepository) Create(Member *model.Member) (*model.Member, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Members[Member.Id()] = &MemberMemory{
		Id:    Member.Id(),
		First: Member.First(),
		Email: Member.Email(),
	}
	return Member, nil
}

func (r *MemberRepository) Read() ([]*model.Member, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]*model.Member, len(r.Members))
	i := 0
	for _, Member := range r.Members {
		users[i] = model.Create(Member.Id, Member.First, Member.Email)
		i++
	}
	return users, nil
}

func (r *MemberRepository) ReadById(id string) (*model.Member, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, Member := range r.Members {
		if Member.Id == id {
			return model.Create(Member.Id, Member.First, Member.Email), nil
		}
	}
	// https://golangbot.com/custom-errors/
	return nil, errors.New(fmt.Sprintf("did not find member with id: %s", id))
}

// ReadByEmail returns the Member if there is a Club Member with this email, an error otherwise
func (r *MemberRepository) ReadByEmail(email string) (*model.Member, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, Member := range r.Members {
		if Member.Email == email {
			return model.Create(Member.Id, Member.First, Member.Email), nil
		}
	}
	return nil, errors.New(fmt.Sprintf("did not find member with email: %s", email))
}

// Empty all members; used for testing
func (r *MemberRepository) Empty() {
	r.mu.Lock()
	defer r.mu.Unlock()

	for k := range r.Members {
		delete(r.Members, k)
	}
}
