package handler

import "github.com/gin-gonic/gin"

type MoviesController struct {
}

func NewMoviesController() *MoviesController {
	return &MoviesController{}
}

func (h MoviesController) SendHello(c *gin.Context) {
	c.JSON(200, "Helloword")
}
