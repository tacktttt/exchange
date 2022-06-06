package order

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/order"
	"github.com/tacktttt/exchange/domain/repository"
	"github.com/tacktttt/exchange/domain/service"
)

type CreateOrderParam struct {
	pairID           uuid.UUID
	amount           int
	position         string
	keyAmount        int
	settlementAmount int
}

type IOrderApplication interface {
	CreateOrder(params CreateOrderParam) error
}

type CreateOrderApplication struct {
	ctx             context.Context
	con             *sql.DB
	pairRepository  repository.IPairRepository
	contractService service.ContractService
}

func NewCreateOrderApplication(
	ctx context.Context,
	con *sql.DB,
	contractService service.ContractService,
) *CreateOrderApplication {
	return &CreateOrderApplication{
		ctx:             ctx,
		con:             con,
		contractService: contractService,
	}
}

func (o *CreateOrderApplication) CreateOrder(params CreateOrderParam) error {
	pair, err := o.pairRepository.GetPairByID(o.con, params.pairID)
	if err != nil {
		return err
	}

	order := order.NewOrder(params.amount, params.position, pair)
	return o.contractService.CreateOrder(order)
}
