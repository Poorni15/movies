package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"movie/internal/models"

	"github.com/google/uuid"
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

func (movieRepository *MovieRepository) GetAll() ([]models.Movie, error) {
	query := `
        SELECT id, imdb_code, title, description, release_year, genre, rating
        FROM movies
    `
	rows, err := movieRepository.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		var idStr string
		if err := rows.Scan(&idStr, &m.ImdbCode, &m.Title, &m.Description, &m.ReleaseYear, &m.Genre, &m.Rating); err != nil {
			return nil, err
		}
		m.Id, err = uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
