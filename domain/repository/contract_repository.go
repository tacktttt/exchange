package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/contract"
)

type IContractRepository interface {
	Create(con *sql.DB, contract *contract.Contract) (*contract.Contract, error)
	GetContractsByUserID(con *sql.DB, userID uuid.UUID) ([]*contract.Contract, error)
}
