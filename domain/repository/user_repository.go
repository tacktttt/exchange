package repository

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/user"
)

type IUserRepository interface {
	GetUserByID(userID uuid.UUID) (*user.IUser, error)
}
