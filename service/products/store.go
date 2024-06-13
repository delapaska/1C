package products

import (
	"database/sql"
	"fmt"

	"github.com/delapaska/1C/models"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

/*
order_id INT REFERENCES orders(id),
Name VARCHAR(255) NOT NULL,
Quantity int  NOT NULL,
Unit VARCHAR(15)
*/
func (s *Store) AddProduct(product models.ProductPayload, order_id int) error {
	query := fmt.Sprintf("INSERT INTO products (name, quantity, unit, order_id) VALUES('%s', '%d', '%s', '%d')", product.Name, product.Quantity, product.Unit, order_id)
	fmt.Println(query)
	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}
	return nil

}

func (s *Store) GetProductsById(orderID int) ([]models.Product, error) {
	query := "SELECT id, name, quantity, unit, order_id FROM products WHERE order_id = $1"
	rows, err := s.db.Query(query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Unit, &product.Order_id); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
