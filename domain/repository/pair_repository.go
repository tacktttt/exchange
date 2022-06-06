package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/pair"
)

type IPairRepository interface {
	GetPairs(con *sql.DB) ([]*pair.Pair, error)
	GetPairsByKeyCurrencyID(con *sql.DB, keyCurrencyID uuid.UUID) ([]*pair.Pair, error)
	GetPairsBySettlementCurrencyID(con *sql.DB, settlementCurrencyID uuid.UUID) ([]*pair.Pair, error)
	GetPairByID(con *sql.DB, pairID uuid.UUID) (*pair.Pair, error)
}
