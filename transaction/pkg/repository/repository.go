package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/repository/sqlite"
)

type Repository struct {
	Withdraw sqlite.Withdraw
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Withdraw: sqlite.NewWithdrawSQLite(db),
	}
}
