package models

type OrderStore interface {
	CreateOrder(OrderPayload) error
}

type ProductStore interface {
	AddProduct(ProductPayload, int) error
	GetProductsById(orderID int) ([]Product, error)
}

type Order struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type OrderPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
type Product struct {
	ID       int    `json:"id"`
	Order_id int    `json:"order_id"`
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
	Unit     string `json:"unit"`
}

type ProductPayload struct {
	Name     string `json:"name" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
	Unit     string `json:"unit"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
