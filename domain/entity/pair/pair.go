package pair

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/currency"
)

type Pair struct {
	ID                 uuid.UUID
	keyCurrency        currency.ICurrency
	settlementCurrency currency.ICurrency
}

func NewPair(
	id uuid.UUID,
	keyCurrency currency.ICurrency,
	settlementCurrency currency.ICurrency,
) *Pair {
	return &Pair{
		ID:                 id,
		keyCurrency:        keyCurrency,
		settlementCurrency: settlementCurrency,
	}
}

func (p *Pair) KeyCurrency() currency.IKeyCurrency {
	return p.keyCurrency
}

func (p *Pair) SettlementCurrency() currency.ISettlementCurrency {
	return p.settlementCurrency
}
