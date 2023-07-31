package repository

import (
	"context"
	"fmt"
	"time"

	"grey/internal/domain"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

type ProductInterface interface {
	CreateProduct(ctx context.Context, input domain.ProductInput) error
	ListProduct(ctx context.Context) ([]domain.Product, error)
	GetProduct(ctx context.Context, productId int) (domain.Product, error)
	UpdatePrice(ctx context.Context, productId int, input domain.ProductPriceInput) error

}

func (r *ProductRepository) CreateProduct(ctx context.Context, input domain.ProductInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	timeNow := time.Now()

	var productId int
	queryProduct := fmt.Sprintf("INSERT INTO %s (name, description, quantity, price_date) values ($1, $2, $3, $4) RETURNING id", tableProduct)
	row := tx.QueryRow(
		queryProduct, 
		input.Name, input.Description, input.Quantity, timeNow,
	)
	err = row.Scan(&productId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	var priceId int
	queryPrice := fmt.Sprintf("INSERT INTO %s (product_id, price, date) values ($1, $2, $3) RETURNING id", tablePrice)
	row = tx.QueryRow(
		queryPrice, 
		productId, input.Price, timeNow,
	)
	err = row.Scan(&priceId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	for _, s := range input.Tags{
		var tagId int
		queryTag := fmt.Sprintf("INSERT INTO %s (name) values ($1) ON CONFLICT (name) DO UPDATE SET name = EXCLUDED.name RETURNING id", tableTag)
		row = tx.QueryRow(
			queryTag, 
			s,
		)
		err = row.Scan(&tagId)
		if err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("")
			return err
		}

		queryProductTag := fmt.Sprintf("INSERT INTO %s (product_id, tag_id) values ($1, $2) RETURNING 0", tableProductTag)
		row = tx.QueryRow(
			queryProductTag, 
			productId, tagId,
		)
		err = row.Scan(&tagId)
		if err != nil {
			tx.Rollback()
			log.Error().Err(err).Msg("")
			return err
		}
	}

	tx.Commit()

	return nil
}


func (r *ProductRepository) ListProduct(ctx context.Context) ([]domain.Product, error) {
	var res []domain.Product

// SELECT prd.*, prc.price, ARRAY_AGG(tg.id) FROM grey.product AS prd JOIN grey.price AS prc ON prd.id=prc.product_id JOIN grey.product_tag AS prd_tg ON prd.id=prd_tg.product_id JOIN grey.tag AS tg ON prd_tg.tag_id=tg.id GROUP BY prd.id, prc.price
// SELECT product_id, MAX (id) FROM grey.price GROUP BY product_id;
// SELECT prc.product_id, prc.price, t.mx FROM (SELECT product_id, MAX(id) AS mx FROM grey.price GROUP BY product_id) t JOIN grey.price prc ON prc.product_id = t.product_id AND t.mx=prc.id;

// with version
// SELECT prd.*, prc.price, ARRAY_AGG(tg.id) FROM grey.product AS prd JOIN grey.price AS prc ON prd.id=prc.product_id AND prd.price_version=prc.version JOIN grey.product_tag AS prd_tg ON prd.id=prd_tg.product_id JOIN grey.tag AS tg ON prd_tg.tag_id=tg.id GROUP BY prd.id, prc.price

	query := fmt.Sprintf(
		// "SELECT prd.id, prd.name, prd.description, prd.quantity, prc.price, ARRAY_AGG(tg.name) FROM %s AS prd JOIN %s AS prc ON prd.id=prc.product_id JOIN %s AS prd_tg ON prd.id=prd_tg.product_id JOIN %s AS tg ON tg.id=prd_tg.tag_id",
		// "SELECT prd.id, prd.name, prd.description, prd.quantity, prc.price, ARRAY_AGG(tg.name) AS tags FROM %s AS prd JOIN %s AS prc ON prd.id=prc.product_id JOIN %s AS prd_tg ON prd.id=prd_tg.product_id JOIN %s AS tg ON prd_tg.tag_id=tg.id GROUP BY prd.id, prc.price",
		// "SELECT prd.id, prd.name, prd.description, prd.quantity, prc.price, ARRAY_AGG(tg.name) AS tags FROM %s AS prd JOIN %s AS prc ON prd.id=prc.product_id JOIN %s AS prd_tg ON prd.id=prd_tg.product_id JOIN %s AS tg ON prd_tg.tag_id=tg.id GROUP BY prd.id, prc.price",
		"SELECT prd.id, prd.name, prd.description, prd.quantity, prc.price, ARRAY_AGG(tg.name) AS tags FROM %s AS prd JOIN %s AS prc ON prd.id=prc.product_id AND prd.price_date=prc.date JOIN %s AS prd_tg ON prd.id=prd_tg.product_id JOIN %s AS tg ON prd_tg.tag_id=tg.id GROUP BY prd.id, prc.price",
		tableProduct, tablePrice, tableProductTag, tableTag,
	)
	err := r.db.Select(&res, query)
	if err != nil{
		log.Error().Err(err).Msg("")
	}

	return res, err
}

func (r *ProductRepository) GetProduct(ctx context.Context, productId int) (domain.Product, error) {
	var res domain.Product

// SELECT prd.*, prc.price, ARRAY_AGG(tg.id) FROM grey.product AS prd JOIN grey.price AS prc ON prd.id=prc.product_id JOIN grey.product_tag AS prd_tg ON prd.id=prd_tg.product_id JOIN grey.tag AS tg ON prd_tg.tag_id=tg.id WHERE prd.id= 18 GROUP BY prd.id, prc.price, prc.id ORDER BY prc.id DESC LIMIT 1

	query := fmt.Sprintf(
		"SELECT prd.id, prd.name, prd.description, prd.quantity, prc.price, ARRAY_AGG(tg.name) AS tags FROM %s AS prd JOIN %s AS prc ON prd.id=prc.product_id AND prd.price_date=prc.date JOIN %s AS prd_tg ON prd.id=prd_tg.product_id JOIN %s AS tg ON prd_tg.tag_id=tg.id WHERE prd.id = %d GROUP BY prd.id, prc.price",
		tableProduct, tablePrice, tableProductTag, tableTag, productId,
	)
	err := r.db.Get(&res, query)
	if err != nil{
		log.Error().Err(err).Msg("")
	}

	return res, err
}

func (r *ProductRepository) UpdatePrice(ctx context.Context, productId int, input domain.ProductPriceInput) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	timeNow := time.Now()
	
	queryProduct := fmt.Sprintf("UPDATE %s SET price_date=$1 WHERE id=$2", tableProduct)
	_, err = tx.Exec(
		queryProduct, 
		timeNow, productId,
	)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	var priceId int
	queryPrice := fmt.Sprintf("INSERT INTO %s (product_id, price, date) values ($1, $2, $3) RETURNING id", tablePrice)
	row := tx.QueryRow(
		queryPrice, 
		productId, input.Price, timeNow,
	)
	err = row.Scan(&priceId)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msg("")
		return err
	}

	tx.Commit()

	return nil
}
