package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/tacktttt/exchange/domain/entity/order"
)

type IOrderRepository interface {
	Create(con *sql.DB, order *order.Order) (*order.Order, error)
	Delete(con *sql.DB, order *order.Order) error
	GetOrdersByPairID(con *sql.DB, pairID uuid.UUID) ([]*order.Order, error)
	GetOrderByID(con *sql.DB, orderID uuid.UUID) (*order.Order, error)
	GetOppositeOrders(con *sql.DB, pairID uuid.UUID, position string, amount int) ([]*order.Order, error)
}
