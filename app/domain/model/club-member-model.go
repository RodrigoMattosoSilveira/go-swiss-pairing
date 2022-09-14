package model

type ClubMember struct {
	id    string
	first string
	email string
}

func NewClubMember(id string, first string, email string) *ClubMember {
	return &ClubMember{
		id:    id,
		first: first,
		email: email,
	}
}

func (cm ClubMember) Id() string {
	return cm.id
}

func (cm ClubMember) First() string {
	return cm.first
}

func (cm ClubMember) Email() string {
	return cm.email
}
