package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	name string
}

func NewUser(
	id uuid.UUID,
	name string,
) *User {
	return &User{
		ID:   id,
		name: name,
	}
}

func (u *User) GetName() string {
	return u.name
}
