package model

import "time"

type Launch struct {
	Id        int       `json:"id" db:"id"`
	Command   string    `json:"command" db:"command"`
	StartTime time.Time `json:"start_time" db:"start_time"`
	EndTime   time.Time `json:"end_time" db:"end_time"`
	Result    string    `json:"result" db:"result"`
	EventId   int       `json:"event_id" db:"event_id"`
}
