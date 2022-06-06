package contract

import (
	"github.com/tacktttt/exchange/domain/entity/execution"
	"github.com/tacktttt/exchange/domain/entity/order"
)

type IContract interface {
	Order() order.IOrder
	OppositeOrders() []order.IOrder
	ExecContract() []execution.IExecution
}
