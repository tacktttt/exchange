package wallet

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/currency"
	"github.com/tacktttt/exchange/domain/entity/user"
)

type Wallet struct {
	ID       uuid.UUID
	amount   int
	user     *user.User
	currency *currency.Currency
}

func NewWallet(
	id uuid.UUID,
	amount int,
	user *user.User,
	currency *currency.Currency,
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

func (k *Wallet) User() *user.User {
	return k.user
}

func (k *Wallet) Currency() *currency.Currency {
	return k.currency
}
