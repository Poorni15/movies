package service

import (
	"encoding/json"
	"fmt"
	"movie/internal/models"
	"net/http"

	"github.com/google/uuid"
)

type OMDBResponse struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Genre      string `json:"Genre"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
	Actors     string `json:"Actors"`
	Response   string `json:"Response"`
	Error      string `json:"Error"`
}

const OMDB_API_KEY = "979be668"

func FetchMovieFromOMDb(imdbCode string) (*models.Movie, error) {
	var OMDBResponse OMDBResponse
	url := fmt.Sprintf("https://www.omdbapi.com/?i=%s&plot=full&apikey=%s", imdbCode, OMDB_API_KEY)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(res.Body).Decode(&OMDBResponse); err != nil || OMDBResponse.Response != "True" {
		return nil, fmt.Errorf("OMDb error: %s", OMDBResponse.Error)
	}
	releaseYear := 0
	fmt.Sscanf(OMDBResponse.Year, "%d", &releaseYear)
	rating := 0
	fmt.Sscanf(OMDBResponse.ImdbRating, "%d", &rating)
	return &models.Movie{Id: uuid.New(), ImdbCode: OMDBResponse.ImdbID, Title: OMDBResponse.Title, Description: OMDBResponse.Title, ReleaseYear: releaseYear, Genre: OMDBResponse.Genre, Rating: rating}, nil
}
