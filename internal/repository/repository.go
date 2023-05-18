package repository

import (
	"github.com/jmoiron/sqlx"
)

type EventRepository struct {
	event
}

type LaunchRepository struct {
	launch
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{NewApiPostgres(db)}
}

func NewLaunchRepository(db *sqlx.DB) *LaunchRepository {
	return &LaunchRepository{NewApiPostgres(db)}
}
