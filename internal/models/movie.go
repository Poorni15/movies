package models

type Movie struct {
	Id          string `json:"id"`
	ImdbCode    string `json:"imdb_code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseYear int    `json:"release_year"`
	Genre       string `json:"genre"`
	Rating      int    `json:"rating"`
}
