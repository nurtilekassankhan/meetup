package handler

import (
	"context"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/model"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/service/accountingsrv"
)

type Handler interface {
	RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error
	ConfirmWithdraw(ctx context.Context, otp string) error
}

type handler struct {
	accounting accountingsrv.Service
}

func (h *handler) RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error {
	return h.accounting.RequestWithdraw(ctx, data)
}

func (h *handler) ConfirmWithdraw(ctx context.Context, otp string) error {
	return h.accounting.ConfirmWithdraw(ctx, otp)
}

func New(accounting accountingsrv.Service) (Handler, error) {
	return &handler{
		accounting: accounting,
	}, nil
}
