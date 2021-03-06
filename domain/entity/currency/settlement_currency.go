package currency

import (
	"github.com/google/uuid"
)

type SettlementCurrency struct {
	ID           uuid.UUID
	symbol       string
	issuedAmount int
}

func LoadSettlementCurrency(
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

func (s *SettlementCurrency) Symbol() string {
	return s.symbol
}

func (s *SettlementCurrency) IssuedAmount() int {
	return s.issuedAmount
}
