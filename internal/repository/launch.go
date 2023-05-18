package repository

import (
	"fmt"
	"watchAndRun/internal/model"
)

type Launch interface {
	AddLaunch(launch model.Launch) error
	GetAllLaunches() ([]model.Launch, error)
}

const launchTable = "launch"

func (r launchRepository) AddLaunch(launch model.Launch) error {
	addLaunch := fmt.Sprintf("INSERT INTO %s (event_id, command, start_time, end_time, result) values ($1, $2, $3, $4, $5)", launchTable)
	_, err := r.db.Exec(addLaunch, launch.EventId, launch.Command, launch.StartTime, launch.EndTime, launch.Result)
	if err != nil {
		return err
	}

	return nil
}

func (r launchRepository) GetAllLaunches() ([]model.Launch, error) {
	var launches []model.Launch

	query := fmt.Sprintf("SELECT * FROM %s", launchTable)
	err := r.db.Select(&launches, query)

	return launches, err
}
