package handler

import (
	"fmt"
	"movie/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MoviesController struct {
	MovieService *service.MovieService
}

func NewMoviesController(service *service.MovieService) *MoviesController {
	return &MoviesController{MovieService: service}
}

func (h MoviesController) SendHello(c *gin.Context) {
	c.JSON(200, "Helloword")
}

func (movieController *MoviesController) Search(c *gin.Context) {
	imdbCode := c.Query("imdb_code")
	title := c.Query("title")
	fmt.Printf("Received search request for imdb_code: %s\n", imdbCode)
	fmt.Printf("Received search request for title: %s\n", title)
	movie, err := movieController.MovieService.Search(title, imdbCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, movie)
}
