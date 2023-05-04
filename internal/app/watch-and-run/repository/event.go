package repository

import (
	"fmt"
	"watchAndRun/internal/app/watch-and-run/model"
)

type event interface {
	AddEvent(event model.Event) error
	GetAllEvents() ([]model.Event, error)
}

func (r ApiPostgres) AddEvent(event model.Event) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	addEvent := fmt.Sprintf("INSERT INTO %s (path, file_name, type, time) values ($1, $2, $3, $4)", r.dbTables.EventTable)
	_, err = tx.Exec(addEvent, event.Path, event.FileName, event.EventType, event.Time)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r ApiPostgres) GetAllEvents() ([]model.Event, error) {
	var events []model.Event

	query := fmt.Sprintf("SELECT * FROM %s", r.dbTables.EventTable)
	err := r.db.Select(&events, query)

	return events, err
}
