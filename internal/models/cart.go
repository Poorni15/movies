package models

import (
	"github.com/google/uuid"
)

type Cart struct {
	ID     uuid.UUID  `json:"id"`
	UserID uuid.UUID  `json:"user_id"`
	Items  []CartItem `json:"items"`
}
