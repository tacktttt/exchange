package wallet

import (
	"time"

	"github.com/tacktttt/exchange/domain/entity/pair"
)

type IOrder interface {
	Amount() int
	Pair() pair.IPair
	KeyAmount() int
	SettlementAmount() int
	UtcCreatedAt() time.Time
}
