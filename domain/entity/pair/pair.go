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

func (p *Pair) GetKeyCurrency() currency.ICurrency {
	return p.keyCurrency
}

func (p *Pair) GetSettlementCurrency() currency.ICurrency {
	return p.settlementCurrency
}
