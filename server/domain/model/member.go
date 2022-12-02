package model

type Member struct {
	id       string
	first    string
	last     string
	email    string
	password string
	cell     string
	rating   int32
	isActive bool
	imageUrl string
}

func Create(
	id string,
	first string,
	last string,
	email string,
	password string,
	cell string,
	rating int32,
	isActive bool,
	imageUrl string,
) *Member {
	return &Member{
		id:       id,
		first:    first,
		last:     last,
		email:    email,
		password: password,
		cell:     cell,
		rating:   rating,
		isActive: isActive,
		imageUrl: imageUrl,
	}
}

func (cm Member) Id() string {
	return cm.id
}

func (cm Member) First() string {
	return cm.first
}

func (cm Member) Last() string {
	return cm.last
}

func (cm Member) Email() string {
	return cm.email
}

func (cm Member) Password() string {
	return cm.password
}

func (cm Member) Cell() string {
	return cm.cell
}

func (cm Member) Rating() int32 {
	return cm.rating
}

func (cm Member) IsActive() bool {
	return cm.isActive
}

func (cm Member) ImageUrl() string {
	return cm.imageUrl
}

func (cm Member) Empty() {}
