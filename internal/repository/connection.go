package repository

import (
	"database/sql"
)

func Connect() (*sql.DB, error) {
	return sql.Open("postgres", "host=localhost port=5432 user=postgres password=poorni1512 dbname=postgres sslmode=disable")
}
