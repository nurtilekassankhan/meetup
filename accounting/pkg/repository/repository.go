package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/repository/sqlite"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/repository/transactionconn"
)

type Repository struct {
	WithdrawRequest sqlite.WithdrawRequest
	Transaction     transactionconn.Connection
}

func NewRepository(db *sqlx.DB, conn transactionconn.Connection) *Repository {
	return &Repository{
		WithdrawRequest: sqlite.NewWithdrawRequestSQLite(db),
		Transaction:     conn,
	}
}
