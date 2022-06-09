package order

import (
	"time"

	"github.com/tacktttt/exchange/domain/entity/execution"
	"github.com/tacktttt/exchange/domain/entity/pair"
)

type IOrder interface {
	Pair() pair.IPair
	Position() string
	Amount() int
	SettlementAmount() int
	Executions() []*execution.Execution
	UtcCreatedAt() time.Time
}
