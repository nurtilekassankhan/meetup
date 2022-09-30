package transaction_errors

import "errors"

var (
	InvalidTransactionId = errors.New("invalid transaction id")
)
