package main

import (
	"fmt"
	"log"
	"movie/internal/handler"
	"movie/internal/repository"
	"movie/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting the application...")
	router := gin.Default()
	db, err := repository.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMoviesService(movieRepository)
	moviesController := handler.NewMoviesController(movieService)
	router.GET("/helloworld", moviesController.SendHello)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
