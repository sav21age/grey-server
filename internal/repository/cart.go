package repository

import (
	"context"
	"errors"
	"fmt"

	"grey/internal/domain"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type CartRepository struct {
	db *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) *CartRepository {
	return &CartRepository{db: db}
}

type CartInterface interface {
	AddProduct(ctx context.Context, userId int, input domain.CartInput) error
	ListCart(ctx context.Context, userId int) (domain.Cart, error)
	CheckoutCart(ctx context.Context, userId int) error
}

func (r *CartRepository) AddProduct(ctx context.Context, userId int, input domain.CartInput) error {
	var priceId, productQuantity int

	queryPriceIdProductQuantity := fmt.Sprintf(
		"SELECT prc.id, prd.quantity FROM %s AS prd JOIN %s AS prc ON prd.id=prc.product_id AND prd.price_date=prc.date WHERE prd.id=$1",
		tableProduct, tablePrice,
	)
	row := r.db.QueryRow(queryPriceIdProductQuantity, input.ProductID)
	err := row.Scan(&priceId, &productQuantity)
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	if productQuantity < input.Quantity{
		return errors.New("not enough of product")		
	}

	tx, err := r.db.Begin()
	if err != nil {
		log.Error().Err(err).Msg("")
		return err
	}

	var cartId int
	queryCart := fmt.Sprintf(
		"INSERT INTO %s (user_id) values ($1) ON CONFLICT (user_id) DO UPDATE SET user_id = EXCLUDED.user_id RETURNING id", 
		tableCart,
	)
	row = tx.QueryRow(queryCart, userId)
	err = row.Scan(&cartId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	var cartItemId int
	queryPrice := fmt.Sprintf(
		"INSERT INTO %s (product_id, cart_id, price_id, quantity) values ($1, $2, $3, $4) RETURNING 0", 
		tableCartItem,
	)
	row = tx.QueryRow(
		queryPrice, 
		input.ProductID, cartId, priceId, input.Quantity,
	)
	err = row.Scan(&cartItemId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	tx.Commit()

	return nil
}

func (r *CartRepository) ListCart(ctx context.Context, userId int) (domain.Cart, error) {
	var cartItem []domain.CartItem
	var total *int

// SELECT prd.name, prc.price, ci.quantity FROM grey.cart_item AS ci JOIN grey.product AS prd ON prd.id=ci.product_id JOIN grey.price AS prc ON ci.price_id=prc.id JOIN grey.cart AS c ON c.id=ci.cart_id WHERE c.user_id=1

	query := fmt.Sprintf(
		"SELECT prd.name, prc.price, ci.quantity, prc.price * ci.quantity AS subtotal FROM %s AS ci JOIN %s AS prd ON prd.id=ci.product_id JOIN %s AS prc ON ci.price_id=prc.id JOIN grey.cart AS c ON c.id=ci.cart_id WHERE c.user_id=$1",
		tableCartItem, tableProduct, tablePrice,
	)
	err := r.db.Select(&cartItem, query, userId)
	if err != nil {
		log.Error().Err(err).Msg("")
		return domain.Cart{}, err
	}

	queryTotal := fmt.Sprintf(
		"SELECT SUM (prc.price * ci.quantity) AS total FROM %s AS ci JOIN %s AS prd ON prd.id=ci.product_id JOIN %s AS prc ON ci.price_id=prc.id JOIN grey.cart AS c ON c.id=ci.cart_id WHERE c.user_id=$1",
		tableCartItem, tableProduct, tablePrice,
	)
	err = r.db.Get(&total, queryTotal, userId)
	if err != nil {
		log.Error().Err(err).Msg("")
		return domain.Cart{}, err
	}

	res := domain.Cart{
		Items: cartItem,
		Total: total,
	}
	
	return res, err
}

func (r *CartRepository) CheckoutCart(ctx context.Context, userId int) error {
// INSERT INTO grey.order (user_id) SELECT c.user_id FROM grey.cart AS c WHERE c.user_id=1 RETURNING id;

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var cartId int
	queryCart := fmt.Sprintf("SELECT id FROM %s WHERE user_id=$1", tableCart)
	row := tx.QueryRow(queryCart, userId)
	err = row.Scan(&cartId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	var orderId int
	queryOrder := fmt.Sprintf(
		"INSERT INTO %s (user_id) SELECT c.user_id FROM %s AS c WHERE c.user_id=$1 RETURNING id", 
		tableOrder, tableCart,
	)
	row = tx.QueryRow(queryOrder, userId)
	err = row.Scan(&orderId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

// INSERT INTO grey.order_item (product_id, order_id, price_id, quantity) SELECT ci.product_id, 7 AS order_id, ci.price_id, ci.quantity FROM grey.cart_item AS ci JOIN grey.cart AS c ON c.id=ci.cart_id WHERE c.user_id=1;
// SELECT product_id, 7 AS order_id, price_id, quantity FROM grey.cart_item WHERE cart_id=3
	queryOrderItem := fmt.Sprintf(
		"INSERT INTO %s (product_id, order_id, price_id, quantity) SELECT product_id, $1 AS order_id, price_id, quantity FROM %s WHERE cart_id=$2", 
		tableOrderItem, tableCartItem,
	)
	_, err = tx.Exec(queryOrderItem, orderId, cartId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}


// UPDATE grey.product AS prd SET quantity=quantity-ci.q FROM (SELECT ci.product_id, ci.quantity AS q FROM grey.cart_item AS ci JOIN grey.cart AS c ON c.id=ci.cart_id WHERE c.user_id=1) AS ci WHERE prd.id=ci.product_id
// SELECT product_id, quantity AS q FROM grey.cart_item WHERE cart_id=3
	queryCartItem := fmt.Sprintf(
		"UPDATE %s AS prd SET quantity=quantity-ci.q FROM (SELECT product_id, quantity AS q FROM %s WHERE cart_id=$1) AS ci WHERE prd.id=ci.product_id", 
		tableProduct, tableCartItem,
	)
	_, err = tx.Exec(queryCartItem, cartId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	queryCartItemDelete := fmt.Sprintf("DELETE FROM %s WHERE cart_id=$1", tableCartItem)
	_, err = tx.Exec(queryCartItemDelete, cartId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	queryCartDelete := fmt.Sprintf("DELETE FROM %s WHERE id=$1", tableCart)
	_, err = tx.Exec(queryCartDelete, cartId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	tx.Commit()

	return nil

}