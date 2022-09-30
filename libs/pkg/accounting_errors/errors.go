package accounting_errors

import "errors"

var (
	InvalidAmount   = errors.New("invalid amount")
	InvalidCurrency = errors.New("invalid currency")
	InvalidOtpCode  = errors.New("invalid otp code")
	ExpiredOtpCode  = errors.New("expired otp code")
)
