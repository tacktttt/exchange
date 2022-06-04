package currency

import (
	"github.com/google/uuid"
)

type SettlementCurrency struct {
	ID           uuid.UUID
	Symbol       string
	IssuedAmount int
}

func NewSettlementCurrency(
	id uuid.UUID,
	symbol string,
	issuedAmount int,
) *SettlementCurrency {
	return &SettlementCurrency{
		ID:           id,
		Symbol:       symbol,
		IssuedAmount: issuedAmount,
	}
}

func (k *SettlementCurrency) GetSymbol() string {
	return k.Symbol
}

func (k *SettlementCurrency) GetIssuedAmount() int {
	return k.IssuedAmount
}
