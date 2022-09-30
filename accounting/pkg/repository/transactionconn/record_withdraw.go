package transactionconn

import (
	"context"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/model"
	pb "github.com/nurtilekassankhan/meetup/accounting/pkg/proto/transaction"
	"time"
)

func (c *connection) RecordWithdraw(ctx context.Context, data *model.WithdrawRequestDatabase) error {

	_, err := c.conn.RecordWithdraw(ctx, &pb.RecordWithdrawRequest{
		Currency:   data.Currency,
		Chain:      data.Blockchain,
		Amount:     data.Amount,
		Address:    data.Address,
		ExecutedAt: time.Now().UTC().String(),
	})

	if err != nil {
		return err
	}

	return err

}
