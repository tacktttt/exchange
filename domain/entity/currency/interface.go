package currency

type ICurrency interface {
	GetSymbol() string
	GetIssuedAmount() int
}
