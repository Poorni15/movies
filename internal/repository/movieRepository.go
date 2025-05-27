package repository

import (
	"database/sql"
	"errors"
	"movie/internal/models"
)

var ErrNotFound = errors.New("movie not found")

type MovieRepository struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{DB: db}
}

func (movierepository *MovieRepository) FindByTitle(title string) (*models.Movie, error) {
	var movie models.Movie
	err := movierepository.DB.QueryRow("SELECT * from movies WHERE title= $1", title).
		Scan(&movie.Id, &movie.ImdbCode, &movie.Title, &movie.Description, &movie.ReleaseYear,
			&movie.Genre, &movie.Rating)
	if err != nil {
		return nil, ErrNotFound
	}
	return &movie, nil
}

func (movieRepository *MovieRepository) Insert(movie *models.Movie) error {
	_, err := movieRepository.DB.Exec(`INSERT INTO MOVIES (imdb_code,title,description,release_year,genre,rating)
	VALUES ($1,$2,$3,$4,$5,$6)`, movie.ImdbCode, movie.Title, movie.Description, movie.ReleaseYear, movie.Genre, movie.Rating)
	return err
}
