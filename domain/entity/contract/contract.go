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
	isContractExecuted bool,
) *Contract {
	return &Contract{
		ID:                 id,
		order:              order,
		oppositeOrders:     oppositeOrders,
		executions:         executions,
		isContractExecuted: isContractExecuted,
	}
}

func (k *Contract) ExecContract() error {
	if k.isContractExecuted {
		return errors.New("contract already executed.")
	}

	switch k.order.Position() {
	case "BUY":
		remainingAmount := k.order.RemainingAmount()
		oppositeOrders := k.OppositeOrders()

		// sortをどこでするか、position毎に取得するか

		for _, oppositeOrder := range oppositeOrders {
			if remainingAmount == 0 {
				break
			}

			if k.order.SettlementAmount() < oppositeOrder.SettlementAmount() {
				break
			}

		}
	case "SELL":
		// SELL
	default:
		return errors.New("invalid position.")
	}

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
