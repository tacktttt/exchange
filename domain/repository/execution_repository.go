package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/execution"
)

type IExecutionRepository interface {
	Create(con *sql.DB, order *execution.Execution) (*execution.Execution, error)
	GetExecutionsByContractID(con *sql.DB, contractID uuid.UUID) ([]*execution.Execution, error)
}
