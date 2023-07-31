package service

import (
	"context"
	"grey/internal/domain"
	"grey/internal/repository"
)

type OrderService struct {
	r   repository.OrderInterface
}

func NewOrderService(r repository.OrderInterface) *OrderService {
	return &OrderService{
		r: r,
	}
}

//go:generate mockgen -source=order.go -destination=mock/order.go

type OrderInterface interface {
	ListOrder(ctx context.Context, userId int) (domain.OrderList, error)
	DetailOrder(ctx context.Context, userId, orderId int) (domain.Order, error)
}

func (s *OrderService) ListOrder(ctx context.Context, userId int) (domain.OrderList, error) {
	res, err := s.r.ListOrder(ctx, userId)
	if err != nil {
		return domain.OrderList{}, err
	}

	return res, nil
}

func (s *OrderService) DetailOrder(ctx context.Context, userId, orderId int) (domain.Order, error) {
	res, err := s.r.DetailOrder(ctx, userId, orderId)
	if err != nil {
		return domain.Order{}, err
	}

	return res, nil
}
