package domain

import "github.com/lib/pq"

// type Product struct {
// 	ID          int    `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// 	Quantity    int    `json:"quantity"`
// 	Price       int    `json:"price"`
// 	Tags        sql.NullString `json:"tags"`
// 	Tags        *string `json:"tags"`
// 	Tags pq.StringArray `json:"tags"`
// }

type Product struct {
	ID int `json:"id"`
	ProductInput
}

type ProductInput struct {
	Name        string         `json:"name" binding:"required"`
	Description string         `json:"description,omitempty"`
	Quantity    int            `json:"quantity" binding:"required,min=1"`
	Price       int            `json:"price" binding:"required,min=1"`
	Tags        pq.StringArray `json:"tags,omitempty" swaggertype:"array,string"`
}

type ProductPriceInput struct {
	Price int `json:"price" binding:"required,min=1"`
}
