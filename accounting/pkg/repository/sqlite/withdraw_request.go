package sqlite

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/model"
)

const createWithdrawRequestsTableSQL = `
	CREATE TABLE IF NOT EXISTS withdraw_requests (
	    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	    currency TEXT NOT NULL,
	    blockchain TEXT NOT NULL,
	    amount TEXT NOT NULL,
	    address TEXT NOT NULL,
	    otp TEXT NOT NULL,
	    created_at DATE NOT NULL
	);
`

type WithdrawRequest interface {
	Create(data *model.WithdrawRequestDatabase) (int, error)
	GetByOtpCode(otp string) (model.WithdrawRequestDatabase, error)
}

type WithdrawRequestSQLite struct {
	db *sqlx.DB
}

func NewWithdrawRequestSQLite(db *sqlx.DB) *WithdrawRequestSQLite {
	return &WithdrawRequestSQLite{db: db}
}

func (r *WithdrawRequestSQLite) Create(data *model.WithdrawRequestDatabase) (int, error) {
	query := fmt.Sprintf("insert into %s (currency, blockchain, amount, address, otp, created_at) values ($1, $2, $3, $4, $5, $6)", withdrawRequestsTable)
	res, err := r.db.Exec(query, data.Currency, data.Blockchain, data.Amount, data.Address, data.Otp, data.CreatedAt)
	if err != nil {
		return 0, nil
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func (r *WithdrawRequestSQLite) GetByOtpCode(otp string) (model.WithdrawRequestDatabase, error) {
	var data model.WithdrawRequestDatabase
	query := fmt.Sprintf("select * FROM %s WHERE otp = $1", withdrawRequestsTable)
	if err := r.db.Get(&data, query, otp); err != nil {
		return data, err
	}
	return data, nil
}

func createWithdrawRequestsTable(db *sqlx.DB) error {
	table, err := db.Prepare(createWithdrawRequestsTableSQL)
	if err != nil {
		return err
	}
	_, err = table.Exec()
	if err != nil {
		return err
	}
	return nil
}
