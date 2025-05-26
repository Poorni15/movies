package repository

import (
	"database/sql"
	"movie/internal/models"

	"github.com/google/uuid"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepository *UserRepository) Create(name string) (*models.User, error) {
	id := uuid.New()
	_, err := userRepository.DB.Exec("INSERT INTO users (id,name) VALUES ($1,$2)", id, name)
	if err != nil {
		return nil, err
	}
	return &models.User{ID: id, Name: name}, nil
}
