package contract

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/execution"
	"github.com/tacktttt/exchange/domain/entity/order"
)

type Contract struct {
	ID             uuid.UUID
	order          *order.Order
	oppositeOrders []*order.Order
	executions     []*execution.Execution
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
		ID:             id,
		order:          order,
		oppositeOrders: oppositeOrders,
		executions:     executions,
	}
}

func (k *Contract) Order() *order.Order {
	return k.order
}

func (k *Contract) OppositeOrders() []*order.Order {
	return k.oppositeOrders
}

func (k *Contract) Executions() []*execution.Execution {
	// TODO
	return k.executions
}
