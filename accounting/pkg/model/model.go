package model

import "time"

type WithdrawRequest struct {
	Currency   string `json:"currency"`
	Blockchain string `json:"blockchain"`
	Amount     string `json:"amount"`
	Address    string `json:"address"`
}

type WithdrawRequestDatabase struct {
	Id         int       `db:"id"`
	Currency   string    `db:"currency"`
	Blockchain string    `db:"blockchain"`
	Amount     string    `db:"amount"`
	Address    string    `db:"address"`
	Otp        string    `db:"otp"`
	CreatedAt  time.Time `db:"created_at"`
}

type WithdrawConfirm struct {
	OtpCode string `json:"otp_code"`
}
