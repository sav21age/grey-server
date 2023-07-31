package domain

import "time"

type Order struct {
	Items []OrderItem `json:"items"`
	Total *int        `json:"total"`
}

type OrderItem struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	SubTotal int    `json:"subtotal"`
}

type OrderList struct {
	Orders []OrderListItem `json:"orders"`
	Total  *int            `json:"total"`
}

type OrderListItem struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
}
