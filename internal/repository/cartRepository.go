package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type CartRepository struct {
	DB *sql.DB
}

func NewCartRepository(db *sql.DB) *CartRepository {
	return &CartRepository{DB: db}
}
func (cartRepository *CartRepository) CreateCart(userID uuid.UUID) (uuid.UUID, error) {
	cartID := uuid.New()
	_, err := cartRepository.DB.Exec("INSERT INTO carts (id, user_id) VALUES ($1, $2)", cartID, userID)
	return cartID, err
}
