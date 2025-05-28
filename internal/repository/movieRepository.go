package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"movie/internal/models"
)

var ErrNotFound = errors.New("movie not found")

type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (movieRepository *MovieRepository) Search(title, imdbCode string) (*models.Movie, error) {

	if title != "" {
		movie, err := movieRepository.findByTitle(title)
		if err != nil {
			return nil, err
		}
		return movie, nil
	}

	if imdbCode != "" {
		movie, err := movieRepository.findByImdbCode(imdbCode)
		if err != nil {
			return nil, err
		}
		return movie, nil
	}
	return nil, errors.New("either title or imdbCode must be provided")
}

func (movierepository *MovieRepository) findByTitle(title string) (*models.Movie, error) {
	var movie models.Movie
	err := movierepository.DB.QueryRow("SELECT * from movies WHERE title= $1", title).
		Scan(&movie.Id, &movie.ImdbCode, &movie.Title, &movie.Description, &movie.ReleaseYear,
			&movie.Genre, &movie.Rating)
	if err != nil {
		return nil, ErrNotFound
	}
	return &movie, nil
}

func (movierepository *MovieRepository) findByImdbCode(imdbCode string) (*models.Movie, error) {
	var movie models.Movie
	err := movierepository.DB.QueryRow("SELECT * from movies WHERE imdb_code= $1", imdbCode).
		Scan(&movie.Id, &movie.ImdbCode, &movie.Title, &movie.Description, &movie.ReleaseYear,
			&movie.Genre, &movie.Rating)
	if err != nil {
		return nil, ErrNotFound
	}
	return &movie, nil
}

func (movieRepository *MovieRepository) Insert(movie *models.Movie) error {
	fmt.Printf("Url formed: %s\n", movie.ImdbCode)
	_, err := movieRepository.DB.Exec(`INSERT INTO MOVIES (imdb_code,title,description,release_year,genre,rating)
	VALUES ($1,$2,$3,$4,$5,$6)`, movie.ImdbCode, movie.Title, movie.Description, movie.ReleaseYear, movie.Genre, movie.Rating)
	return err
}
