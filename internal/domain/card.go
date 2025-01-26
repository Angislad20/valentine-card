package domain

import "time"

type Card struct {
	ID       int
	Message  string
	AudioURL string
	Theme    string
	CreateAt time.Time
}
