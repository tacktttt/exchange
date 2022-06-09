package service

import (
	"context"
	"database/sql"

	"github.com/tacktttt/exchange/domain/entity/contract"
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
	contractRepository  repository.IContractRepository
	executionRepository repository.IExecutionRepository
}

func NewContractService(
	ctx context.Context,
	con *sql.DB,
	orderRepository repository.IOrderRepository,
	contractRepository repository.IContractRepository,
	executionRepository repository.IExecutionRepository,
) *ContractService {
	return &ContractService{
		ctx:                 ctx,
		con:                 con,
		orderRepository:     orderRepository,
		contractRepository:  contractRepository,
		executionRepository: executionRepository,
	}
}

func (c *ContractService) CreateOrder(order *order.Order) error {
	orders, err := c.orderRepository.GetOrdersByPositionAndAmount(c.con, order.Position(), order.Amount())
	if err != nil {
		return err
	}

	return c.createContract(order, orders)
}

func (c *ContractService) createContract(order *order.Order, oppositeOrders []*order.Order) error {
	contract := contract.NewContract(order, oppositeOrders)

	err := contract.ExecContract()
	if err != nil {
		return err
	}

	_, err = c.contractRepository.Create(c.con, contract)
	if err != nil {
		return err
	}

	for _, execution := range contract.Executions() {
		_, err := c.executionRepository.Create(c.con, execution)
		if err != nil {
			return err
		}
	}

	return nil
}
