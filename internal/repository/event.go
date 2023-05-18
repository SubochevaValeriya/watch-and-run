package repository

import (
	"fmt"
	"watchAndRun/internal/model"
)

type event interface {
	AddEvent(event model.Event) (int64, error)
	GetAllEvents() ([]model.Event, error)
}

const eventTable = "event"

func (r ApiPostgres) AddEvent(event model.Event) (int64, error) {
	tx, err := r.db.Beginx() //no need in tx we can r.db.Exec()
	if err != nil {
		return -1, err
	}

	var id int64
	addEvent := fmt.Sprintf("INSERT INTO %s (path, file_name, type, time) values ($1, $2, $3, $4) RETURNING id", eventTable)
	//remove dbTables and set table name explicit, no need to config, we need migration to rename table after creating
	err = tx.QueryRow(addEvent, event.Path, event.FileName, event.EventType, event.Time).Scan(&id)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	if tx.Commit() != nil {
		return -1, err //return zero val=0, not -1
	}

	return id, nil
}

func (r ApiPostgres) GetAllEvents() ([]model.Event, error) {
	var events []model.Event

	query := fmt.Sprintf("SELECT * FROM %s", eventTable)
	err := r.db.Select(&events, query)

	return events, err
}
