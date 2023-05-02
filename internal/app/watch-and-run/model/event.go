package model

import "time"

type event struct {
	id        int       `json:"id" db:"id"`
	path      string    `json:"path" db:"path"`
	eventType string    `json:"type" db:"type"`
	time      time.Time `json:"time" db:"time"`
}
