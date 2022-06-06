package order

import (
	"time"

	"github.com/tacktttt/exchange/domain/entity/pair"
)

type IOrder interface {
	Amount() int
	Position() string
	Pair() pair.IPair
	KeyAmount() int
	SettlementAmount() int
	UtcCreatedAt() time.Time
}
