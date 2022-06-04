package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/currency"
	"github.com/tacktttt/exchange/domain/entity/pair"
)

func main() {
	keyCurrency := currency.NewKeyCurrency(uuid.MustParse("a3e48008-9947-4323-ab65-a7a97c17e416"), "USD", 100)
	settlementCurrency := currency.NewSettlementCurrency(uuid.MustParse("a3e48008-9947-4323-ab65-a7a97c17e417"), "JPY", 10000)
	pair := pair.NewPair(uuid.MustParse("a3e48008-9947-4323-ab65-a7a97c17e418"), keyCurrency, settlementCurrency)

	fmt.Println(pair)
}
