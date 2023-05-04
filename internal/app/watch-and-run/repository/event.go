package repository

import (
	"fmt"
	"watchAndRun/internal/app/watch-and-run/model"
)

type event interface {
	AddEvent(event model.Event) (int64, error)
	GetAllEvents() ([]model.Event, error)
}

func (r ApiPostgres) AddEvent(event model.Event) (int64, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return -1, err
	}

	addEvent := fmt.Sprintf("INSERT INTO %s (path, file_name, type, time) values ($1, $2, $3, $4)", r.dbTables.EventTable)
	result, err := tx.Exec(addEvent, event.Path, event.FileName, event.EventType, event.Time)

	if err != nil {
		tx.Rollback()
		return -1, err
	}

	if tx.Commit() != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (r ApiPostgres) GetAllEvents() ([]model.Event, error) {
	var events []model.Event

	query := fmt.Sprintf("SELECT * FROM %s", r.dbTables.EventTable)
	err := r.db.Select(&events, query)

	return events, err
}
