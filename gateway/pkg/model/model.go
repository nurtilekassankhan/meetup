package model

type WithdrawRequest struct {
	Currency string `json:"currency"`
	Chain    string `json:"chain"`
	Amount   string `json:"amount"`
	Address  string `json:"address"`
}

type WithdrawConfirm struct {
	OtpCode string `json:"otp_code"`
}
