package handler

import (
	"context"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/model"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/service/transactionsrv"
)

type Handler interface {
	RecordWithdraw(ctx context.Context, data *model.WithdrawDatabase) error
}

type handler struct {
	transaction transactionsrv.Service
}

func New(transaction transactionsrv.Service) (Handler, error) {
	return &handler{
		transaction: transaction,
	}, nil
}
