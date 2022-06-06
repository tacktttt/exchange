package execution

import (
	"github.com/google/uuid"
)

type Execution struct {
	ID     uuid.UUID
	amount int
}

func NewExecution(
	id uuid.UUID,
	amount int,
) *Execution {
	return &Execution{
		ID:     id,
		amount: amount,
	}
}

func (k *Execution) Amount() int {
	return k.amount
}
