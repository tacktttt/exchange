package wallet

import (
	"github.com/tacktttt/exchange/domain/entity/currency"
	"github.com/tacktttt/exchange/domain/entity/user"
)

type IWallet interface {
	Amount() int
	User() user.IUser
	Currency() currency.ICurrency
}
