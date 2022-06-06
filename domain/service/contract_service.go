package service

import (
	"context"
	"database/sql"

	"github.com/tacktttt/exchange/domain/entity/contract"
	"github.com/tacktttt/exchange/domain/entity/execution"
	"github.com/tacktttt/exchange/domain/entity/order"
	"github.com/tacktttt/exchange/domain/repository"
)

type IContractService interface {
	CreateOrder(order order.IOrder) error
}

type ContractService struct {
	ctx                 context.Context
	con                 *sql.DB
	orderRepository     repository.IOrderRepository
	executionRepository repository.IExecutionRepository
}

func NewContractService(
	ctx context.Context,
	con *sql.DB,
	orderRepository repository.IOrderRepository,
	executionRepository repository.IExecutionRepository,
) *ContractService {
	return &ContractService{
		ctx:                 ctx,
		con:                 con,
		orderRepository:     orderRepository,
		executionRepository: executionRepository,
	}
}

func (c *ContractService) CreateOrder(order *order.Order) error {
	switch order.Position() {
	case "BUY":
		return nil
	case "SELL":
		return c.sellOrder(order)
	}

	return nil
}

func (c *ContractService) sellOrder(order *order.Order) error {
	buyOrders, err := c.orderRepository.GetBuyOrdersBySellAmount(c.con, order.Amount())
	if err != nil {
		return err
	}

	contract := contract.NewContract(order, buyOrders)

	// TODO: create contracts

	err = c.createExecutions(contract.Executions())
	if err != nil {
		return err
	}

	return nil
}

func (c *ContractService) createExecutions(executions []*execution.Execution) error {
	for _, execution := range executions {
		_, err := c.executionRepository.Create(c.con, execution)
		if err != nil {
			return err
		}
	}
	return nil
}
