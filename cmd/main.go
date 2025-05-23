package main

import (
	"fmt"
	"movie/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running main")
	router := gin.Default()
	moviesController := handler.NewMoviesController()
	router.GET("/helloworld", moviesController.SendHello)
	router.Run(":8080")
}
