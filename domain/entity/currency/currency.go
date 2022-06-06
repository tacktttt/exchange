package currency

import (
	"github.com/google/uuid"
)

type Currency struct {
	ID           uuid.UUID
	symbol       string
	issuedAmount int
}

func NewCurrency(
	id uuid.UUID,
	symbol string,
	issuedAmount int,
) *Currency {
	return &Currency{
		ID:           id,
		symbol:       symbol,
		issuedAmount: issuedAmount,
	}
}

func (k *Currency) Symbol() string {
	return k.symbol
}

func (k *Currency) IssuedAmount() int {
	return k.issuedAmount
}
