package models

import "github.com/google/uuid"

type CartItem struct {
	Id      uuid.UUID `json:"id"`
	CartId  uuid.UUID `json:"cart_id"`
	MovieId uuid.UUID `json:"movie_id"`
}
