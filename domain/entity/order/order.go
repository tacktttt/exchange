package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/pair"
)

type Order struct {
	ID               uuid.UUID
	amount           int
	position         string
	pair             *pair.Pair
	keyAmount        int
	settlementAmount int
	utcCreatedAt     time.Time
}

func NewOrder(
	amount int,
	position string,
	pair *pair.Pair,
) *Order {
	return &Order{
		amount:   amount,
		position: position,
		pair:     pair,
	}
}

func LoadOrder(
	id uuid.UUID,
	amount int,
	position string,
	pair *pair.Pair,
) *Order {
	return &Order{
		ID:       id,
		amount:   amount,
		position: position,
		pair:     pair,
	}
}

func (k *Order) Amount() int {
	return k.amount
}

func (k *Order) Position() string {
	return k.position
}

func (k *Order) Pair() *pair.Pair {
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
