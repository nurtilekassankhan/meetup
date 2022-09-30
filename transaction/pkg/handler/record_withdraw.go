package handler

import (
	"context"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/model"
)

func (h *handler) RecordWithdraw(ctx context.Context, data *model.WithdrawDatabase) error {
	return h.transaction.RecordWithdraw(ctx, data)
}
