package repository

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/execution"
)

type IExecutionRepository interface {
	Create(order *execution.IExecution) (*execution.IExecution, error)
	GetExecutionsByContractID(contractID uuid.UUID) ([]*execution.IExecution, error)
}
