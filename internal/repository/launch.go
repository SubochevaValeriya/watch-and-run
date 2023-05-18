package repository

import (
	"fmt"
	"watchAndRun/internal/model"
)

type launch interface {
	AddLaunch(launch model.Launch) error
	GetAllLaunches() ([]model.Launch, error)
}

const launchTable = "launch"

func (r ApiPostgres) AddLaunch(launch model.Launch) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	addLaunch := fmt.Sprintf("INSERT INTO %s (event_id, command, start_time, end_time, result) values ($1, $2, $3, $4, $5)", launchTable)
	_, err = tx.Exec(addLaunch, launch.EventId, launch.Command, launch.StartTime, launch.EndTime, launch.Result)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r ApiPostgres) GetAllLaunches() ([]model.Launch, error) {
	var launches []model.Launch

	query := fmt.Sprintf("SELECT * FROM %s", launchTable)
	err := r.db.Select(&launches, query)

	return launches, err
}