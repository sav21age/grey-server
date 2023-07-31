package domain

type Cart struct {
	Items []CartItem `json:"items"`
	Total *int       `json:"total"`
}

type CartItem struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	SubTotal int    `json:"subtotal"`
}

type CartInput struct {
	ProductID int `json:"product_id" binding:"required"`
	Quantity  int `json:"quantity" binding:"required,min=1"`
}
