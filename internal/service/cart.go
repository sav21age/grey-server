package service

import (
	"context"
	"grey/internal/domain"
	"grey/internal/repository"
)

type CartService struct {
	r   repository.CartInterface
}

func NewCartService(r repository.CartInterface) *CartService {
	return &CartService{
		r:   r,
	}
}

//go:generate mockgen -source=cart.go -destination=mock/cart.go

type CartInterface interface {
	AddProduct(ctx context.Context, userId int, input domain.CartInput) error
	ListCart(ctx context.Context, userId int) (domain.Cart, error)
	CheckoutCart(ctx context.Context, userId int) error
}

func (s *CartService) AddProduct(ctx context.Context, userId int, input domain.CartInput) error {
	if err := s.r.AddProduct(ctx, userId, input); err != nil {
		return err
	}

	return nil
}

func (s *CartService) ListCart(ctx context.Context, userId int) (domain.Cart, error) {
	res, err := s.r.ListCart(ctx, userId)
	if err != nil {
		return domain.Cart{}, err
	}

	return res, nil
}

func (s *CartService) CheckoutCart(ctx context.Context, userId int) error {
	err := s.r.CheckoutCart(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}