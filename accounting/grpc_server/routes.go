package grpc_server

import (
	"context"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/model"
	pb "github.com/nurtilekassankhan/meetup/accounting/pkg/proto/accounting"
)

func (s *server) RequestWithdraw(ctx context.Context, in *pb.RequestWithdrawRequest) (*pb.RequestWithdrawResponse, error) {
	if err := s.handlr.RequestWithdraw(ctx, &model.WithdrawRequest{
		Currency:   in.GetCurrency(),
		Blockchain: in.GetChain(),
		Amount:     in.GetAmount(),
		Address:    in.GetAddress(),
	}); err != nil {
		return nil, err
	}
	return &pb.RequestWithdrawResponse{}, nil
}

func (s *server) ConfirmWithdraw(ctx context.Context, in *pb.ConfirmWithdrawRequest) (*pb.ConfirmWithdrawResponse, error) {
	if err := s.handlr.ConfirmWithdraw(ctx, in.GetCode()); err != nil {
		return nil, err
	}
	return &pb.ConfirmWithdrawResponse{}, nil
}
