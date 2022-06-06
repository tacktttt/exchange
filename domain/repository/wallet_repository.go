package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/wallet"
)

type IWalletRepository interface {
	GetWallets(con *sql.DB, userID uuid.UUID) ([]*wallet.Wallet, error)
	GetWalletByID(con *sql.DB, userID uuid.UUID, walletID uuid.UUID) (*wallet.Wallet, error)
}
