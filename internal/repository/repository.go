package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	tableUser       = "grey.user"
	tableProduct    = "grey.product"
	tableTag        = "grey.tag"
	tablePrice      = "grey.price"
	tableProductTag = "grey.product_tag"
	tableCart       = "grey.cart"
	tableCartItem   = "grey.cart_item"
	tableOrder      = "grey.order"
	tableOrderItem  = "grey.order_item"
)

type Repository struct {
	User    UserInterface
	Product ProductInterface
	Cart    CartInterface
	Order   OrderInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserRepository(db),
		Product: NewProductRepository(db),
		Cart:    NewCartRepository(db),
		Order:   NewOrderRepository(db),
	}
}
