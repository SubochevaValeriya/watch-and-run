package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	event
	launch
}

type repository struct {
	Repository
}

func NewRepository(db *sqlx.DB, dbTables DbTables) *repository {
	return &repository{NewApiPostgres(db, dbTables)}
}
