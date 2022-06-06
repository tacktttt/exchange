package repository

import (
	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/order"
)

type IOrderRepository interface {
	Create(order *order.IOrder) (*order.IOrder, error)
	Delete(order *order.IOrder) error
	GetOrdersByPairID(pairID uuid.UUID) ([]*order.IOrder, error)
	GetOrderByID(orderID uuid.UUID) (*order.IOrder, error)
}
