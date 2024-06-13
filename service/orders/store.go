package orders

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

func (s *Store) CreateOrder(order models.OrderPayload) error {
	query := fmt.Sprintf("INSERT INTO orders (name, description) VALUES('%s', '%s')", order.Name, order.Description)
	fmt.Println(query)
	_, err := s.db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}
