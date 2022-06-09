package repository

import (
	"database/sql"

	"github.com/tacktttt/exchange/domain/entity/contract"
)

type IContractRepository interface {
	Create(con *sql.DB, contract *contract.Contract) (*contract.Contract, error)
}
