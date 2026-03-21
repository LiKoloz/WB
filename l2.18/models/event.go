package models

import "time"

// Event моделька события
type Event struct {
	Id    int
	Title string
	Text  string
	Time  time.Time
}
