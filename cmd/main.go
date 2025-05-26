package main

import (
	"fmt"
	"io"
	"movie/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Running main")
	router := gin.Default()
	moviesController := handler.NewMoviesController()
	router.GET("/helloworld", moviesController.SendHello)
	router.Run(":8080")

	url := "https://api.themoviedb.org/3/authentication"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI0NDg3N2M4NTMwMTIyNTQ5MDllNGExN2M5NDMyYjgyNSIsIm5iZiI6MTc0ODIzODA0OC4wOSwic3ViIjoiNjgzM2ZlZTAxZjQxYmU1OGI5YjZiNTBlIiwic2NvcGVzIjpbImFwaV9yZWFkIl0sInZlcnNpb24iOjF9.CvbOc7CzW6ct1YOxdyU1N7KwTHQx-7npjKZ0Xr44m9g")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
}
