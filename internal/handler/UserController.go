package handler

import (
	"movie/internal/models"
	"movie/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository *repository.UserRepository
}

func NewUserController(service *repository.UserRepository) *UserController {
	return &UserController{userRepository: service}
}

func (u UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u.userRepository.Create(user.Name)
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}
