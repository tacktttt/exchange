package currency

type ICurrency interface {
	Symbol() string
	IssuedAmount() int
}

type IKeyCurrency interface {
	ICurrency
}

type ISettlementCurrency interface {
	ICurrency
}
