package service

import "github.com/tacktttt/exchange/domain/repository"

type IContractService interface{}

type ContractService struct {
	orderRepository     repository.IOrderRepository
	executionRepository repository.IExecutionRepository
}

func NewContractService(
	orderRepository repository.IOrderRepository,
	executionRepository repository.IExecutionRepository,
) *ContractService {
	return &ContractService{
		orderRepository:     orderRepository,
		executionRepository: executionRepository,
	}
}
