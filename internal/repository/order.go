package repository

import (
	"context"
	"fmt"

	"grey/internal/domain"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

type OrderInterface interface {
	ListOrder(ctx context.Context, userId int) (domain.OrderList, error)
	DetailOrder(ctx context.Context, userId, orderId int) (domain.Order, error)
}

func (r *OrderRepository) ListOrder(ctx context.Context, userId int) (domain.OrderList, error) {
	var orderList []domain.OrderListItem
	var total *int

	queryOrder := fmt.Sprintf("SELECT id, date FROM %s WHERE user_id=$1", tableOrder)
	err := r.db.Select(&orderList, queryOrder, userId)
	if err != nil {
		log.Error().Err(err).Msg("")
		return domain.OrderList{}, err
	}

	queryTotal := fmt.Sprintf("SELECT COUNT(id) AS total FROM %s WHERE user_id=$1", tableOrder)
	err = r.db.Get(&total, queryTotal, userId)
	if err != nil {
		log.Error().Err(err).Msg("")
		return domain.OrderList{}, err
	}

	res := domain.OrderList{
		Orders: orderList,
		Total: total,
	}

	return res, err
}

func (r *OrderRepository) DetailOrder(ctx context.Context, userId, orderId int) (domain.Order, error) {
	var checkOrderId *int
	queryOrderId := fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1 AND id=$2", tableOrder)
	err := r.db.Get(&checkOrderId, queryOrderId, userId, orderId)
	if err != nil {
		log.Error().Err(err).Msg("order not exists")
		return domain.Order{}, err
	}

	var orderItem []domain.OrderItem
	var total *int

	// SELECT prd.name, prc.price, oi.quantity, prc.price * oi.quantity AS subtotal FROM grey.order_item AS oi JOIN grey.product AS prd ON prd.id=oi.product_id JOIN grey.price AS prc ON oi.price_id=prc.id JOIN grey.order AS o ON o.id=oi.order_id WHERE o.user_id=1 AND o.id=11
	queryOrder := fmt.Sprintf(
		"SELECT prd.name, prc.price, oi.quantity, prc.price * oi.quantity AS subtotal FROM %s AS oi JOIN %s AS prd ON prd.id=oi.product_id JOIN %s AS prc ON oi.price_id=prc.id JOIN %s AS o ON o.id=oi.order_id WHERE o.user_id=$1 AND o.id=$2", 
		tableOrderItem, tableProduct, tablePrice, tableOrder,
	)
	err = r.db.Select(&orderItem, queryOrder, userId, orderId)
	if err != nil {
		log.Error().Err(err).Msg("")
		return domain.Order{}, err
	}

	queryTotal := fmt.Sprintf(
		"SELECT SUM (prc.price * oi.quantity) AS total FROM %s AS oi JOIN %s AS prd ON prd.id=oi.product_id JOIN %s AS prc ON oi.price_id=prc.id JOIN %s AS o ON o.id=oi.order_id WHERE o.user_id=$1 AND o.id=$2", 
		tableOrderItem, tableProduct, tablePrice, tableOrder,
	)
	err = r.db.Get(&total, queryTotal, userId, orderId)
	if err != nil {
		log.Error().Err(err).Msg("")
		return domain.Order{}, err
	}

	res := domain.Order{
		Items: orderItem,
		Total: total,
	}

	return res, err
}

