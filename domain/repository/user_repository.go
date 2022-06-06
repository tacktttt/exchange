package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/user"
)

type IUserRepository interface {
	GetUserByID(con *sql.DB, userID uuid.UUID) (*user.User, error)
}
