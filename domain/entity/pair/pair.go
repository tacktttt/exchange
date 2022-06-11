package pair

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/currency"
)

type Pair struct {
	ID                 uuid.UUID
	keyCurrency        *currency.KeyCurrency
	settlementCurrency *currency.SettlementCurrency
}

func LoadPair(
	id uuid.UUID,
	keyCurrency *currency.KeyCurrency,
	settlementCurrency *currency.SettlementCurrency,
) *Pair {
	return &Pair{
		ID:                 id,
		keyCurrency:        keyCurrency,
		settlementCurrency: settlementCurrency,
	}
}

func (p *Pair) KeyCurrency() *currency.KeyCurrency {
	return p.keyCurrency
}

func (p *Pair) SettlementCurrency() *currency.SettlementCurrency {
	return p.settlementCurrency
}
