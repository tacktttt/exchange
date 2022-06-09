package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/order"
	"github.com/tacktttt/exchange/domain/repository"
	"github.com/tacktttt/exchange/domain/service"
)

type CreateOrderParam struct {
	PairID           uuid.UUID
	Position         string
	Amount           int
	SettlementAmount int
}

type ICreateOrderUsecase interface {
	Exec(params CreateOrderParam) error
}

type CreateOrderUsecase struct {
	ctx             context.Context
	con             *sql.DB
	pairRepository  repository.IPairRepository
	contractService service.ContractService
}

func NewCreateOrderUsecase(
	ctx context.Context,
	con *sql.DB,
	contractService service.ContractService,
) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		ctx:             ctx,
		con:             con,
		contractService: contractService,
	}
}

func (o *CreateOrderUsecase) Exec(params CreateOrderParam) error {
	pair, err := o.pairRepository.GetPairByID(o.con, params.PairID)
	if err != nil {
		return err
	}

	order := order.NewOrder(pair, params.Position, params.Amount, params.SettlementAmount)
	return o.contractService.CreateOrder(order)
}
