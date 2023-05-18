package repository

import (
	"github.com/jmoiron/sqlx"
)

type eventRepository struct {
	db *sqlx.DB
}

type launchRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) Event {
	return &eventRepository{db: db}
}

func NewLaunchRepository(db *sqlx.DB) Launch {
	return &launchRepository{db}
}
