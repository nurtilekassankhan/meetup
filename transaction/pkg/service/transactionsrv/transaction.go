package transactionsrv

import (
	"context"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/model"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/repository"
)

type Service interface {
	RecordWithdraw(ctx context.Context, data *model.WithdrawDatabase) error
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) (Service, error) {
	return &service{repo: repo}, nil
}
