package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for the database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{
			DB: db,
		},
	}
}

// Movie is the type for movies
type Movie struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Year int `json:"year"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime int `json:"runtime"`
	Rating int `json:"rating"`
	MPAARating string `json:"mpaa_rating"`
	CreatedAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
	MovieGenre []MovieGenre `json:"genres"`
}

// Genre is the type of genre
type Genre struct {
	ID int `json:"-"`
	GenreName string `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}

// MovieGenre is the type of movie genre
type MovieGenre struct {
	ID int `json:"-"`
	MovieID int `json:"-"`
	GenreID int `json:"-"`
	Genre Genre `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}