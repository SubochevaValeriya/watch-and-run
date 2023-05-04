package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repo interface {
	event
	launch
}

type Repository struct {
	Repo
}

func NewRepository(db *sqlx.DB, dbTables DbTables) *Repository {
	return &Repository{NewApiPostgres(db, dbTables)}
}
