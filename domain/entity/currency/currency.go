package currency

type Currency interface {
	GetSymbol() string
	GetIssuedAmount() int
}
