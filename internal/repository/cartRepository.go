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

func (cartRepository *CartRepository) AddMoviesToCart(cardId, movieId uuid.UUID) (uuid.UUID, error) {
	cartItemID := uuid.New()
	_, error := cartRepository.DB.Exec("INSERT INTO cart_items(id,cart_id,movie_id) VALUES ($1,$2,$3)", cartItemID, cardId, movieId)
	return cartItemID, error
}
func (cartRepository *CartRepository) GetMoviesInCart(cartID uuid.UUID) ([]uuid.UUID, error) {
	rows, err := cartRepository.DB.Query("SELECT movie_id FROM cart_items WHERE cart_id=$1", cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movieIDs []uuid.UUID
	for rows.Next() {
		var id uuid.UUID
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		movieIDs = append(movieIDs, id)
	}
	return movieIDs, nil
}
