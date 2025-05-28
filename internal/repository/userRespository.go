package repository

import (
	"database/sql"
	"movie/internal/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB             *sql.DB
	CartRepository CartRepository
}

func NewUserRepository(db *sql.DB, cartRepository CartRepository) *UserRepository {
	return &UserRepository{DB: db, CartRepository: cartRepository}
}

func (userRepository *UserRepository) Create(name string) (*models.User, error) {
	id := uuid.New()
	_, err := userRepository.DB.Exec("INSERT INTO users (id,name) VALUES ($1,$2)", id, name)
	if err != nil {
		return nil, err
	}
	userRepository.CartRepository.CreateCart(id)
	return &models.User{ID: id, Name: name}, nil
}

func (userRepository *UserRepository) GetByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := userRepository.DB.QueryRow("SELECT id, name FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
