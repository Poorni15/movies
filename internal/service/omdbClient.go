package service

import (
	"encoding/json"
	"fmt"
	"movie/internal/models"
	"net/http"
	"net/url"

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

func FetchMovieFromOMDb(imdbCode, title string) (*models.Movie, error) {
	fmt.Printf("imdb coed: %s\n", imdbCode)
	fmt.Printf("title : %s\n", title)
	var OMDBResponse OMDBResponse
	baseURL := "http://www.omdbapi.com/"
	fmt.Printf("base url: %s\n", baseURL)
	params := url.Values{}
	if imdbCode != "" {
		params.Add("i", imdbCode)
	} else {
		params.Add("t", title)
	}
	params.Add("apikey", OMDB_API_KEY)
	endpoint := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Printf("Url formed: %s\n", endpoint)
	res, err := http.Get(endpoint)
	fmt.Printf("response body: %#v\n", res.Body)
	if err != nil {
		return nil, err
	}
	if err := json.NewDecoder(res.Body).Decode(&OMDBResponse); err != nil || OMDBResponse.Response != "True" {
		return nil, fmt.Errorf("OMDb error: %s", OMDBResponse.Error)
	}
	fmt.Printf("response body: %#v\n", OMDBResponse)
	releaseYear := 0
	fmt.Sscanf(OMDBResponse.Year, "%d", &releaseYear)
	rating := 0
	fmt.Sscanf(OMDBResponse.ImdbRating, "%d", &rating)
	return &models.Movie{Id: uuid.New(), ImdbCode: OMDBResponse.ImdbID, Title: OMDBResponse.Title, Description: OMDBResponse.Title, ReleaseYear: releaseYear, Genre: OMDBResponse.Genre, Rating: rating}, nil
}
