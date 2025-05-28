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

func (cartRepository *CartRepository) AddMoviesToCart(cartID, movieID uuid.UUID) (uuid.UUID, error) {
	cartItemID := uuid.New()
	_, err := cartRepository.DB.Exec(
		"INSERT INTO cart_items (id, cart_id, movie_id) VALUES ($1, $2, $3)",
		cartItemID, cartID, movieID,
	)
	if err != nil {
		return uuid.Nil, err
	}
	return cartItemID, nil
}
func (cartRepository *CartRepository) GetOrCreateCart(userID uuid.UUID) (uuid.UUID, error) {
	var cartID uuid.UUID
	err := cartRepository.DB.QueryRow("SELECT id FROM carts WHERE user_id = $1", userID).Scan(&cartID)
	if err != nil {
		if err == sql.ErrNoRows {
			return cartRepository.CreateCart(userID)
		}
		return uuid.Nil, err
	}
	return cartID, nil
}
func (r *CartRepository) GetMoviesInCart(cartID uuid.UUID) ([]string, error) {
	rows, err := r.DB.Query(`
        SELECT m.title
        FROM cart_items ci
        INNER JOIN movies m ON ci.movie_id = m.id
        WHERE ci.cart_id = $1
    `, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var titles []string
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			return nil, err
		}
		titles = append(titles, title)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return titles, nil
}
