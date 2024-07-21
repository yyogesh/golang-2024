package models

import "time"

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e) // events = [...events, newEvent]
}

func GetEvents() []Event {
	return events
}
