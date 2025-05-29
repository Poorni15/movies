package handler

import (
	"fmt"
	"movie/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartController struct {
	cartRepository *repository.CartRepository
}

func NewCartController(cartRepository *repository.CartRepository) *CartController {
	return &CartController{cartRepository: cartRepository}
}

func (cc *CartController) ViewCart(c *gin.Context) {
	userIDStr := c.Param("user_id")
	fmt.Printf("Received search request for user_id: %s\n", userIDStr)
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	cartID, err := cc.cartRepository.GetOrCreateCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve or create cart"})
		return
	}

	fmt.Printf("Cart ID: %s\n", cartID)

	movieTitles, err := cc.cartRepository.GetMoviesInCart(cartID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movies in cart"})
		return
	}

	fmt.Printf("Movies in cart: %v\n", movieTitles)

	c.JSON(http.StatusOK, gin.H{
		"cart_id": cartID,
		"movies":  movieTitles,
	})
}

func (cc *CartController) AddToCart(c *gin.Context) {
	var input struct {
		UserID  uuid.UUID `json:"user_id"`
		MovieID uuid.UUID `json:"movie_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	cartID, err := cc.cartRepository.GetOrCreateCart(input.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve or create cart"})
		return
	}

	cartItemID, err := cc.cartRepository.AddMoviesToCart(cartID, input.MovieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add movie to cart"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"cart_item_id": cartItemID})
}
