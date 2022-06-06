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
	GetBuyOrdersBySellAmount(con *sql.DB, sellAmount int) ([]*order.Order, error)
	GetSellOrdersByBuyAmount(con *sql.DB, buyAmount int) ([]*order.Order, error)
}
