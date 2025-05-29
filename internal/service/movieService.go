package service

import (
	"fmt"
	"movie/internal/models"
	"movie/internal/repository"
)

type MovieService struct {
	MovieRepository *repository.MovieRepository
}

func NewMoviesService(repository *repository.MovieRepository) *MovieService {
	return &MovieService{MovieRepository: repository}
}

func (movieService *MovieService) Search(title, imdbCode string) (*models.Movie, error) {
	fmt.Println("Inside movie service")
	var movie *models.Movie
	var err error
	movie, err = movieService.MovieRepository.Search(title, imdbCode)
	if err == nil {
		return movie, nil
	}
	if imdbCode != "" || title != "" {
		movie, err = FetchMovieFromOMDb(imdbCode, title)
		fmt.Printf("response body: %#v\n", movie)
		if err != nil {
			return nil, err
		}
		err = movieService.MovieRepository.Insert(movie)
		if err != nil {
			return nil, err
		}
		return movie, nil
	}

	return nil, repository.ErrNotFound
}

func (movieService *MovieService) List() ([]models.Movie, error) {
	return movieService.MovieRepository.GetAll()
}
