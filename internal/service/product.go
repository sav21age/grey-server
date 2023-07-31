package service

import (
	"context"
	"grey/internal/domain"
	"grey/internal/repository"
)

type ProductService struct {
	r repository.ProductInterface
}

func NewProductService(r repository.ProductInterface) *ProductService {
	return &ProductService{
		r: r,
	}
}

//go:generate mockgen -source=product.go -destination=mock/product.go

type ProductInterface interface {
	CreateProduct(ctx context.Context, input domain.ProductInput) error
	ListProduct(ctx context.Context) ([]domain.Product, error)
	GetProduct(ctx context.Context, productId int) (domain.Product, error)
	UpdatePrice(ctx context.Context, productId int, input domain.ProductPriceInput) error
}

func (s *ProductService) CreateProduct(ctx context.Context, input domain.ProductInput) error {
	if err := s.r.CreateProduct(ctx, input); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) ListProduct(ctx context.Context) ([]domain.Product, error) {
	res, err := s.r.ListProduct(ctx)
	if err != nil {
		return []domain.Product{}, err
	}

	return res, nil
}

func (s *ProductService) GetProduct(ctx context.Context, productId int) (domain.Product, error) {
	res, err := s.r.GetProduct(ctx, productId)
	if err != nil {
		return domain.Product{}, err
	}

	return res, nil
}

func (s *ProductService) UpdatePrice(ctx context.Context, productId int, input domain.ProductPriceInput) error {
	if err := s.r.UpdatePrice(ctx, productId, input); err != nil {
		return err
	}

	return nil
}