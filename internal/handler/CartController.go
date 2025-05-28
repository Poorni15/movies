package handler

import "movie/internal/repository"

type CartController struct {
	cartRepository *repository.CartRepository
}

func NewCartController(cartRepository *repository.CartRepository) *CartController {
	return &CartController{cartRepository: cartRepository}
}
