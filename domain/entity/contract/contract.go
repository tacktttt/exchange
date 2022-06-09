package contract

import (
	"errors"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/execution"
	"github.com/tacktttt/exchange/domain/entity/order"
)

type Contract struct {
	ID                 uuid.UUID
	order              *order.Order
	oppositeOrders     []*order.Order
	executions         []*execution.Execution
	isContractExecuted bool
}

func NewContract(
	order *order.Order,
	oppositeOrders []*order.Order,
) *Contract {
	return &Contract{
		order:          order,
		oppositeOrders: oppositeOrders,
	}
}

func LoadContract(
	id uuid.UUID,
	order *order.Order,
	oppositeOrders []*order.Order,
	executions []*execution.Execution,
) *Contract {
	return &Contract{
		ID:                 id,
		order:              order,
		oppositeOrders:     oppositeOrders,
		executions:         executions,
		isContractExecuted: true,
	}
}

func (k *Contract) ExecContract() error {
	if k.isContractExecuted {
		return errors.New("contract already executed.")
	}

	// TODO execute contract
	k.isContractExecuted = true

	return nil
}

func (k *Contract) Order() *order.Order {
	return k.order
}

func (k *Contract) OppositeOrders() []*order.Order {
	return k.oppositeOrders
}

func (k *Contract) Executions() []*execution.Execution {
	return k.executions
}

func (k *Contract) IsContractExecuted() bool {
	return k.isContractExecuted
}
