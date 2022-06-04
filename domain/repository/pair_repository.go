package repository

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/pair"
)

type IPairRepository interface {
	GetPairs() ([]*pair.IPair, error)
	GetPairsByKeyCurrencyID(keyCurrencyID uuid.UUID) ([]*pair.IPair, error)
	GetPairsBySettlementCurrencyID(settlementCurrencyID uuid.UUID) ([]*pair.IPair, error)
	GetPairByID(pairID uuid.UUID) (*pair.IPair, error)
}
