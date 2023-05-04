package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, fmt.Errorf("connect error: %s", err)
	}

	return db, nil
}

type ApiPostgres struct {
	db       *sqlx.DB
	dbTables DbTables
}

type DbTables struct {
	EventTable  string
	LaunchTable string
}

func NewApiPostgres(db *sqlx.DB, dbTables DbTables) *ApiPostgres {
	return &ApiPostgres{db: db,
		dbTables: dbTables}
}
