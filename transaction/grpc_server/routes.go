package grpc_server

import (
	"context"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/model"
	pb "github.com/nurtilekassankhan/meetup/transaction/pkg/proto/transaction"
	"time"
)

func (s *server) RecordWithdraw(ctx context.Context, in *pb.RecordWithdrawRequest) (*pb.RecordWithdrawResponse, error) {

	if err := s.handlr.RecordWithdraw(ctx, &model.WithdrawDatabase{
		Currency:   in.GetCurrency(),
		Blockchain: in.GetChain(),
		Amount:     in.GetAmount(),
		Address:    in.GetAddress(),
		ExecutedAt: time.Now().UTC(),
	}); err != nil {
		return nil, err
	}

	return &pb.RecordWithdrawResponse{}, nil

}
