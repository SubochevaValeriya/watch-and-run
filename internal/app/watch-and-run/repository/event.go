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
	var id int64
	addEvent := fmt.Sprintf("INSERT INTO %s (path, file_name, type, time) values ($1, $2, $3, $4) RETURNING id", r.dbTables.EventTable)
	err = tx.QueryRow(addEvent, event.Path, event.FileName, event.EventType, event.Time).Scan(&id)
	//
	//result, err := tx.QueryRow()
	//tx.Exec(addEvent, event.Path, event.FileName, event.EventType, event.Time)

	if err != nil {
		tx.Rollback()
		return -1, err
	}

	if tx.Commit() != nil {
		return -1, err
	}

	return id, nil
}

func (r ApiPostgres) GetAllEvents() ([]model.Event, error) {
	var events []model.Event

	query := fmt.Sprintf("SELECT * FROM %s", r.dbTables.EventTable)
	err := r.db.Select(&events, query)

	return events, err
}
