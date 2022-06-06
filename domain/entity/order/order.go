package wallet

import (
	"time"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/pair"
)

type Order struct {
	ID               uuid.UUID
	amount           int
	pair             pair.IPair
	keyAmount        int
	settlementAmount int
	utcCreatedAt     time.Time
}

func NewOrder(
	id uuid.UUID,
	amount int,
) *Order {
	return &Order{
		ID:     id,
		amount: amount,
	}
}

func (k *Order) Amount() int {
	return k.amount
}

func (k *Order) Pair() pair.IPair {
	return k.pair
}

func (k *Order) KeyAmount() int {
	return k.keyAmount
}

func (k *Order) SettlementAmount() int {
	return k.settlementAmount
}

func (k *Order) UtcCreatedAt() time.Time {
	return k.utcCreatedAt
}
