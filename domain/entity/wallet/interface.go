package wallet

import (
	"github.com/tacktttt/exchange/domain/entity/currency"
	"github.com/tacktttt/exchange/domain/entity/user"
)

type IWallet interface {
	GetAmount() int
	GetUser() user.IUser
	GetCurrency() currency.ICurrency
}
