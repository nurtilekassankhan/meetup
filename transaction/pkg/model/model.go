package model

import "time"

type WithdrawDatabase struct {
	Id         int       `db:"id"`
	Currency   string    `db:"currency"`
	Blockchain string    `db:"blockchain"`
	Amount     string    `db:"amount"`
	Address    string    `db:"address"`
	ExecutedAt time.Time `db:"executed_at"`
}
