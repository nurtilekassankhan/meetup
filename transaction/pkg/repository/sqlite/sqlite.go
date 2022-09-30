package sqlite

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const (
	withdrawsTable = "withdraws"
)

func NewSQLiteDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "../meetup.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func RunMigrateScripts(db *sqlx.DB) error {
	if err := createWithdrawsTable(db); err != nil {
		return err
	}
	return nil
}
