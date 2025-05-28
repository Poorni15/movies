package main

import (
	"fmt"
	"log"
	"movie/internal/handler"
	"movie/internal/repository"
	"movie/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting the application...")

	router := gin.Default()
	db, err := repository.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMoviesService(movieRepository)
	moviesController := handler.NewMoviesController(movieService)
	userRepository := repository.NewUserRepository(db)
	userController := handler.NewUserController(userRepository)
	router.GET("/helloworld", moviesController.SendHello)
	router.GET("/movies/search", moviesController.Search)
	router.POST("/users", userController.Create)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
