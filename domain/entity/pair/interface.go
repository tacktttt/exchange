package pair

import (
	"github.com/tacktttt/exchange/domain/entity/currency"
)

type IPair interface {
	GetKeyCurrency() currency.ICurrency
	GetSettlementCurrency() currency.ICurrency
}
