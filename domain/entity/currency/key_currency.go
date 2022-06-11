package currency

import (
	"github.com/google/uuid"
)

type KeyCurrency struct {
	ID           uuid.UUID
	symbol       string
	issuedAmount int
}

func LoadKeyCurrency(
	id uuid.UUID,
	symbol string,
	issuedAmount int,
) *KeyCurrency {
	return &KeyCurrency{
		ID:           id,
		symbol:       symbol,
		issuedAmount: issuedAmount,
	}
}

func (k *KeyCurrency) Symbol() string {
	return k.symbol
}

func (k *KeyCurrency) IssuedAmount() int {
	return k.issuedAmount
}

func (k *KeyCurrency) Dummy() int {
	return 0
}
