package accounting_errors

import "errors"

var (
	InvalidAmount = errors.New("invalid amount")
)