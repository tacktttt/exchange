package contract

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/execution"
	"github.com/tacktttt/exchange/domain/entity/order"
)

type Contract struct {
	ID             uuid.UUID
	order          order.IOrder
	oppositeOrders []order.IOrder
	executions     []execution.IExecution
}

func NewContract(
	id uuid.UUID,
	order order.IOrder,
	oppositeOrders []order.IOrder,
) *Contract {
	return &Contract{
		ID:             id,
		order:          order,
		oppositeOrders: oppositeOrders,
	}
}

func (k *Contract) Order() order.IOrder {
	return k.order
}

func (k *Contract) OppositeOrders() []order.IOrder {
	return k.oppositeOrders
}

func (k *Contract) ExecContract() []execution.IExecution {
	// TODO: 約定
	return k.executions
}
