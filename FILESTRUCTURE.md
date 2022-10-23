# File Structure
It will reflect Domain Driven Design principles; I'll base it on Yusuke Hatanaka [Clean Architecture in Go](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1). I'll modify it slightly to reflect some ideas I have about how to evolve the `OO Entities` into a paradigm reflecting the atomic nature of µ-services and a seamless integration of their data models.

# Clean Architecture
`Clean architecture` has some the following layers:
![Conceptual Layer](illustrations/conceptual_layers.png)

There are 4 layers, blue, green, red and yellow layers there in order from the outside:
**_external_** - the blue layer
**_interface_**: the green layer
**_usecase_**: the red layer
**_domain_**: the yellow layer

A key concept is to ensure The most important thing about clean architecture is to make interfaces through each layer.

We will use the `Clean Architecture` model to implement a `Domain Driven Design` architecture using the folder structure discussed in this document.

# The structure
```text
.
├── Makefile
├── README.md
├── app
│   ├── domain
│   │   ├── model
│   │   ├── repository
│   │   └── service
│   ├── interface
│   │   ├── persistence
│   │   └── rpc
│   ├── registry
│   └── usecase
├── cmd
│   └── td
│       └── main.go
└── vendor
├── vendor packages
|...
```

The top directory contains three directories:
- **_app_**: application's packages root directory
- **_cmd_**: application's main package directory
- **_vendor_**: application's vendor packages directory

## Domain - Yellow Layer
Hosted at the `app/domain` folder, a.k.a. `Entity`, it is at the core of the `Clean Architecture`, it holds the `Enterprise Businees Rules` and has three packages:
- **_model_**: has aggregate, entity and value object; perhaps here is where we represent the `Forms`;
- **_repository_**: has repository interfaces of aggregate; perhaps here is where we represent the `Shadows`;
- **_service_**: has application services that depend on several `Shadows`;

### Model
With time the `model` package will grow to aggregate many `Forms` and related `Value Objects`:
```go
package model

type Member struct {
	id string
	first string
	email string
}

func Create (id string, first string, email string) *Member {
	return &Member{
		id:    id,
		first: first,
		email: email
	}
}

func (cm Member)  Id () string {
	return cm.id
}

func (cm Member)  First () string {
	return cm.first
}

func (cm Member)  Email () string {
	return cm.email
}
```

### Repository
They are the `in-memory` data structures used to model our `Forms`. They might, but not always do, mimic the data structures used to persist the system state. We inject an infrastructure layer between the repository and the persistence hardware to translate between them and to enable seamlessly replace hardware persistence. (TODO: need to flush this out):

```go
package repository

import "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/model"

type MemberRepository interface {
	// FindAll find all Member members
	FindAll() ([]*model.Member, error)
	// FindByEmail find a Member member with the given email
	FindByEmail(email string) (*model.Member, error)
	// Save Member member with
	Save(*model.Member) (*model.Member, error)
}
```

****NOTE** that this layer does not know where the `Member Member` `Form` is saved or serialized.

### Service
Here we implement logic to manipulate the models. For exampleto validate that the `Member Member email` is unique we would write something like:
```go
func (u *Member) Duplicated(email string) bool {
        // Find Member member by email from persistence layer...
}
```

The `Duplicated` function is business logic that does not belong in the `Member Member` model. We solve it by ading a service layer like below:
```go
package service

import "github.com/RodrigoMattosoSilveira/go-swiss-pairing/app/domain/repository"

type MemberService struct {
	repo repository. MemberRepository
}

func (s *MemberService) Duplicated(email string) error {
	Member, err := s.repo.FindByEmail(email)
	if Member != nil {
		return fmt.Errorf("%s already exists", email)
	}
	if err != nil {
		return err
	}
	return nil
}
```

## Use Cases - Red Layer
Hosted at the `app/usecase` folder, it holds units of one operation, such as registering and listing `Member members`, as suggested by the following interface:
```go
type UserUsecase interface {
    ListUser() ([]*User, error)
    RegisterUser(email string) error
}
```
 
A more comprehensive `Member Member` implementation would be something like:
```go
type MemberUsecase struct {
    repo    repository.MemberRepository
    service *service.MemberService
}
func NewMemberUsecase(repo repositoryMemberRepository, service *service.MemberService) *MemberUsecase {
    return &ubMemberUsecase {
        repo:    repo,
        service: service,
    }
}
// ListMember list all Member members
func (cm *MemberUsecase) ListMember() ([]*Member, error) {
    Members, err := cm.repo.FindAll()
    if err != nil {
        return nil, err
    }
    return toMember(Members), nil
}
// RegisterMember Register a new Member member
func (cm *MemberUsecase) RegisterMember(first string, email string) error {
    uid, err := uuid.NewRandom()
    if err != nil {
        return err
    }
    if err := u.service.Duplicated(email); err != nil {
        return err
    }
    Member := model.NewMember(uid.String(), first, email)
    if err := u.repo.Save(Member); err != nil {
        return err
    }
    return nil
}
```

`MemberUsercase` depends on two packages, i) repository.MemberRepository interface and ii) *service.MemberService struct. Above, we injected  these two packages when we initialized MemberUsecase. An alternate and more efficient mechanism to set up those dependencies is to use a DI container, which we will discuss shortly.

Also, note that `FindAll` above retrieves `Model.Members`, not `Members` encapsulating undesirable business know how when retrieving Members; hence we expanded `Model.Member` to include logic, `toMember` to return only `Members`:
```go
type Member struct {
    id string
    first string
    email string
}
...
func toMember(Members []*model.Member) []*Member {
    res := make([]*Member, len(Members))
    for i, Member := range Members {
        res[i] = &Member{
            ID:    Member.GetID(),
			First:  Member.First,
            Email: user.GetEmail(),
        }
    }
    return res
}
```

## Interfaces - Green Layer
Hosted at the `app/interface` folder, it holds concrete objects like API endpoint handlers, data repository, and RPC:
- **_persistence_**: Logic to handle the domain data, `Forms`, in memory and persistence; 
- **_rpc_**: API to access `Forms' Shadows`

### Persistence
Hosted at the `app/interface/persistence folder`, below is a concrete implementation of the repository, in memory; we would require a different one to persist it in another medium:
```go
type MemberRepository struct {
    mu    *sync.Mutex
	Members map[string]*Member
}

// NewMemberRepository return the repo
func NewMemberRepository() *MemberRepository {
    return &MemberRepository {
        mu:    &sync.Mutex{},
        Members: map[string]*Member{},
    }
}

// FindAll find all Member members 
func (r *MemberRepository) FindAll() ([]*model.Member, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    Members := make([]*model.Member, len(r.Members))
    i := 0
    for _, Member := range r.Members {
        Members[i] = model.NewMember(Member.ID, Member.First, Member.Email)
        i++
    }
    return Members, nil
}

// FindByEmail find a Member member based on their email
func (r *MemberRepository) FindByEmail(email string) (*model.Member, error) {
    r.mu.Lock()
    defer r.mu.Unlock()
    for _, Member := range r.Members {
        if Member.Email == email {
            return  model.NewMember(Member.ID, Member.First, Member.Email), nil
        }
    }
    return nil, nil
}

// Save Add a new Member Member
func (r *MemberRepository) Save(Member *model.Member) error {
    r.mu.Lock()
    defer r.mu.Unlock()
        r.Members[user.GetID()] = &Member {
        ID:    Member.GetID(),
		First:  Member.GetFirst(), 
        Email: Member.GetEmail(),
    }
    return nil
}
```

Note that, regardless the medium, the `Model.Member` remaining unchanged, remaining oblivious of the actual implementation:
```go
type Member struct {
    ID    string
	First string
    Email string
}
```

### RPC
Hosted at the `app/interface/rpc` folder, we will use gRPC to provide repository access to external sources:
TODO Add this

## Direct Injection
`Dependency injection` (DI) is a software pattern that calls to initialize an object at its creation time, using their already initialized dependencies. We will use [Wire](https://github.com/google/wire) a code generation tool that automates connecting components using dependency injection.
TODO Add this