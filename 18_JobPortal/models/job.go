package models

import "time"

type Job struct {
	ID          int64
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Skills      []string  `binding:"required"`
	Experience  int       `binding:"required"`
	SalaryRange string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64     `binding:"required"`
}
