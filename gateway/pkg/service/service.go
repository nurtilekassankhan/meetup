package service

import (
	"context"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/model"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/repository"
)

type Service interface {
	RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error // gateway -> service
	ConfirmWithdraw(ctx context.Context, data *model.WithdrawConfirm) error // gateway -> service -> service
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{repo: repo}
}

func (s *service) RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error {
	return s.repo.Accounting.RequestWithdraw(ctx, data)
}

func (s *service) ConfirmWithdraw(ctx context.Context, data *model.WithdrawConfirm) error {
	return s.repo.Accounting.ConfirmWithdraw(ctx, data)
}
