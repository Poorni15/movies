package handler

import (
	"movie/internal/models"

	"github.com/gin-gonic/gin"
)

type MoviesController struct {
}

type movieStore interface {
	GetByName(movies *[]models.Movie)
}

func NewMoviesController() *MoviesController {
	return &MoviesController{}
}

func (h MoviesController) SendHello(c *gin.Context) {
	c.JSON(200, "Helloword")
}

func (h MoviesController) SearchByName(c *gin.Context) {

}
