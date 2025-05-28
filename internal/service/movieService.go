package service

import (
	"movie/internal/models"
	"movie/internal/repository"
)

type MovieService struct {
	MovieRepository repository.MovieRepository
}

func NewMoviesController(repository repository.MovieRepository) *MovieService {
	return &MovieService{MovieRepository: repository}
}

func (movieService *MovieService) Search(title, imdbCode string) (*models.Movie, error) {
	var movie *models.Movie
	var err error
	movie, err = movieService.MovieRepository.Search(title, imdbCode)
	if err == nil {
		return movie, nil
	}
	if imdbCode != "" || title != "" {
		movie, err = FetchMovieFromOMDb(imdbCode, title)
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
