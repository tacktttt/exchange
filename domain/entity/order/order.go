package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/execution"
	"github.com/tacktttt/exchange/domain/entity/pair"
)

type Order struct {
	ID               uuid.UUID
	pair             *pair.Pair
	position         string
	amount           int
	settlementAmount int
	executions       []*execution.Execution
	utcCreatedAt     time.Time
}

func NewOrder(
	pair *pair.Pair,
	position string,
	amount int,
	settlementAmount int,
) *Order {
	return &Order{
		pair:             pair,
		position:         position,
		amount:           amount,
		settlementAmount: settlementAmount,
	}
}

func LoadOrder(
	id uuid.UUID,
	pair *pair.Pair,
	position string,
	amount int,
	settlementAmount int,
	executions []*execution.Execution,
	utcCreatedAt time.Time,
) *Order {
	return &Order{
		ID:               id,
		pair:             pair,
		position:         position,
		amount:           amount,
		settlementAmount: settlementAmount,
		executions:       executions,
		utcCreatedAt:     utcCreatedAt,
	}
}

func (k *Order) Pair() *pair.Pair {
	return k.pair
}

func (k *Order) Position() string {
	return k.position
}

func (k *Order) Amount() int {
	return k.amount
}

func (k *Order) AllocatedAmount() int {
	alocatedAmount := 0
	for _, execution := range k.Executions() {
		alocatedAmount += execution.Amount()
	}
	return k.amount
}

func (k *Order) RemainingAmount() int {
	return k.amount - k.AllocatedAmount()
}

func (k *Order) SettlementAmount() int {
	return k.settlementAmount
}

func (k *Order) Executions() []*execution.Execution {
	return k.executions
}

func (k *Order) UtcCreatedAt() time.Time {
	return k.utcCreatedAt
}
