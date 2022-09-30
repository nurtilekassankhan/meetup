package sqlite

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/model"
)

const createWithdrawsTableSQL = `
	CREATE TABLE IF NOT EXISTS withdraws (
	    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	    currency TEXT NOT NULL,
	    blockchain TEXT NOT NULL,
	    amount TEXT NOT NULL,
	    address TEXT NOT NULL,
	    executed_at DATE NOT NULL
	);
`

type Withdraw interface {
	Create(data *model.WithdrawDatabase) (int, error)
}

type WithdrawSQLite struct {
	db *sqlx.DB
}

func NewWithdrawSQLite(db *sqlx.DB) *WithdrawSQLite {
	return &WithdrawSQLite{db: db}
}

func (r *WithdrawSQLite) Create(data *model.WithdrawDatabase) (int, error) {
	query := fmt.Sprintf("insert into %s (currency, blockchain, amount, address, executed_at) values ($1, $2, $3, $4, $5)", withdrawsTable)
	res, err := r.db.Exec(query, data.Currency, data.Blockchain, data.Amount, data.Address, data.ExecutedAt)
	if err != nil {
		return 0, nil
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func createWithdrawsTable(db *sqlx.DB) error {
	table, err := db.Prepare(createWithdrawsTableSQL)
	if err != nil {
		return err
	}
	_, err = table.Exec()
	if err != nil {
		return err
	}
	return nil
}
