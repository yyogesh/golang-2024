package models

import "time"

type Movie struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	RunTime     int       `json:"run_time"`
	Description string    `json:"description"`
	Rating      string    `json:"rating"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"_"`
	UpdatedAt   time.Time `json:"_"`
}
