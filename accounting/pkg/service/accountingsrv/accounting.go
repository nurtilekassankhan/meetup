package accountingsrv

import (
	"context"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/model"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/repository"
)

type Service interface {
	RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error
	ConfirmWithdraw(ctx context.Context, otp string) error
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) (Service, error) {
	return &service{repo: repo}, nil
}
