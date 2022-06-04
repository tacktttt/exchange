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

func (s *SettlementCurrency) GetSymbol() string {
	return s.symbol
}

func (s *SettlementCurrency) GetIssuedAmount() int {
	return s.issuedAmount
}
