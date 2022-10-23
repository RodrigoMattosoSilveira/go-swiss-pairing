package model

type Member struct {
	id    string
	first string
	email string
}

func Create(id string, first string, email string) *Member {
	return &Member{
		id:    id,
		first: first,
		email: email,
	}
}

func (cm Member) Id() string {
	return cm.id
}

func (cm Member) First() string {
	return cm.first
}

func (cm Member) Email() string {
	return cm.email
}
