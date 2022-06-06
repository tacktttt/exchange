package wallet

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/currency"
	"github.com/tacktttt/exchange/domain/entity/user"
)

type Wallet struct {
	ID       uuid.UUID
	amount   int
	user     user.IUser
	currency currency.ICurrency
}

func NewWallet(
	id uuid.UUID,
	amount int,
	user user.IUser,
	currency currency.ICurrency,
) *Wallet {
	return &Wallet{
		ID:       id,
		amount:   amount,
		user:     user,
		currency: currency,
	}
}

func (k *Wallet) Amount() int {
	return k.amount
}

func (k *Wallet) User() user.IUser {
	return k.user
}

func (k *Wallet) Currency() currency.ICurrency {
	return k.currency
}