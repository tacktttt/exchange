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
	oppositeBuyOrders  []*order.Order
	oppositeSellOrders []*order.Order
	executions         []*execution.Execution
	isContractExecuted bool
}

func NewContract(
	order *order.Order,
	oppositeBuyOrders []*order.Order,
	oppositeSellOrders []*order.Order,
) *Contract {
	return &Contract{
		order:              order,
		oppositeBuyOrders:  oppositeBuyOrders,
		oppositeSellOrders: oppositeSellOrders,
	}
}

func LoadContract(
	id uuid.UUID,
	order *order.Order,
	oppositeBuyOrders []*order.Order,
	oppositeSellOrders []*order.Order,
	executions []*execution.Execution,
	isContractExecuted bool,
) *Contract {
	return &Contract{
		ID:                 id,
		order:              order,
		oppositeBuyOrders:  oppositeBuyOrders,
		oppositeSellOrders: oppositeSellOrders,
		executions:         executions,
		isContractExecuted: isContractExecuted,
	}
}

func (k *Contract) ExecContract() error {
	if k.isContractExecuted {
		return errors.New("contract already executed")
	}

	switch k.order.Position() {
	case "BUY":
		remainingAmount := k.order.RemainingAmount()
		sellOrders := k.OppositeSellOrders()

		for _, sellOrder := range sellOrders {
			if remainingAmount == 0 {
				break
			}

			if k.order.SettlementAmount() < sellOrder.SettlementAmount() {
				break
			}

		}
	case "SELL":
		// SELL
	default:
		return errors.New("invalid position")
	}

	k.isContractExecuted = true

	return nil
}

func (k *Contract) Order() *order.Order {
	return k.order
}

func (k *Contract) OppositeSellOrders() []*order.Order {
	return k.oppositeSellOrders
}

func (k *Contract) OppositeBuyOrders() []*order.Order {
	return k.oppositeBuyOrders
}

func (k *Contract) Executions() []*execution.Execution {
	return k.executions
}

func (k *Contract) IsContractExecuted() bool {
	return k.isContractExecuted
}
