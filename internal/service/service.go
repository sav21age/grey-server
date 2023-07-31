package service

import (
	"grey/config"
	"grey/internal/repository"
)

type Service struct {
	User    UserInterface
	Product ProductInterface
	Cart    CartInterface
	Order   OrderInterface
}

func NewService(r *repository.Repository, config *config.Config) *Service {

	return &Service{
		User:    NewUserService(r.User, config),
		Product: NewProductService(r.Product),
		Cart:    NewCartService(r.Cart),
		Order:   NewOrderService(r.Order),
	}
}
