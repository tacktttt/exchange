package currency

import (
	"github.com/google/uuid"
)

type SettlementCurrency struct {
	ID           uuid.UUID
	symbol       string
	issuedAmount int
}

func NewSettlementCurrency(
	id uuid.UUID,
	symbol string,
	issuedAmount int,
) *SettlementCurrency {
	return &SettlementCurrency{
		ID:           id,
		symbol:       symbol,
		issuedAmount: issuedAmount,
	}
}

func (k *SettlementCurrency) GetSymbol() string {
	return k.symbol
}

func (k *SettlementCurrency) GetIssuedAmount() int {
	return k.issuedAmount
}
