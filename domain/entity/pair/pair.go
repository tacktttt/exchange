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

func (k *Pair) GetKeyCurrency() currency.ICurrency {
	return k.keyCurrency
}

func (k *Pair) GetSettlementCurrency() currency.ICurrency {
	return k.settlementCurrency
}
