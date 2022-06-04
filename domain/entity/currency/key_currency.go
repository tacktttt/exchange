package currency

import (
	"github.com/google/uuid"
)

type KeyCurrency struct {
	ID           uuid.UUID
	Symbol       string
	IssuedAmount int
}

func NewKeyCurrency(
	id uuid.UUID,
	symbol string,
	issuedAmount int,
) *KeyCurrency {
	return &KeyCurrency{
		ID:           id,
		Symbol:       symbol,
		IssuedAmount: issuedAmount,
	}
}

func (k *KeyCurrency) GetSymbol() string {
	return k.Symbol
}

func (k *KeyCurrency) GetIssuedAmount() int {
	return k.IssuedAmount
}

func (k *KeyCurrency) Dummy() int {
	return 0
}
