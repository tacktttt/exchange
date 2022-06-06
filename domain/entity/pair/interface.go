package pair

import (
	"github.com/tacktttt/exchange/domain/entity/currency"
)

type IPair interface {
	KeyCurrency() currency.IKeyCurrency
	SettlementCurrency() currency.ISettlementCurrency
}
