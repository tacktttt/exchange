package execution

import (
	"github.com/google/uuid"
)

type Execution struct {
	ID          uuid.UUID
	ContractID  uuid.UUID
	BuyOrderID  uuid.UUID
	SellOrderID uuid.UUID
	amount      int
}

func NewExecution(
	contractID uuid.UUID,
	buyOrderID uuid.UUID,
	sellOrderID uuid.UUID,
	amount int,
) *Execution {
	return &Execution{
		ContractID:  contractID,
		BuyOrderID:  buyOrderID,
		SellOrderID: sellOrderID,
		amount:      amount,
	}
}

func LoadExecution(
	id uuid.UUID,
	contractID uuid.UUID,
	buyOrderID uuid.UUID,
	sellOrderID uuid.UUID,
	amount int,
) *Execution {
	return &Execution{
		ID:          id,
		ContractID:  contractID,
		BuyOrderID:  buyOrderID,
		SellOrderID: sellOrderID,
		amount:      amount,
	}
}

func (k *Execution) Amount() int {
	return k.amount
}
