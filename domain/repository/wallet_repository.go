package repository

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/wallet"
)

type IWalletRepository interface {
	GetWallets(userID uuid.UUID) ([]*wallet.IWallet, error)
	GetWalletByID(userID uuid.UUID, walletID uuid.UUID) (*wallet.IWallet, error)
}
