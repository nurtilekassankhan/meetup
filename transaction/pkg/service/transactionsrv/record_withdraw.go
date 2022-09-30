package transactionsrv

import (
	"context"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/model"
)

func (s *service) RecordWithdraw(ctx context.Context, data *model.WithdrawDatabase) error {

	_, err := s.repo.Withdraw.Create(data)
	if err != nil {
		return err
	}

	return nil

}
