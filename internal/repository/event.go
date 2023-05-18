package repository

import (
	"fmt"
	"watchAndRun/internal/model"
)

type Event interface {
	AddEvent(event model.Event) (int64, error)
	GetAllEvents() ([]model.Event, error)
}

const eventTable = "event"

func (r eventRepository) AddEvent(event model.Event) (int64, error) {
	var id int64
	addEvent := fmt.Sprintf("INSERT INTO %s (path, file_name, type, time) values ($1, $2, $3, $4) RETURNING id", eventTable)
	err := r.db.QueryRow(addEvent, event.Path, event.FileName, event.EventType, event.Time).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r eventRepository) GetAllEvents() ([]model.Event, error) {
	var events []model.Event

	query := fmt.Sprintf("SELECT * FROM %s", eventTable)
	err := r.db.Select(&events, query)

	return events, err
}
