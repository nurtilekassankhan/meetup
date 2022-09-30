package accountingconn

import (
	"context"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/model"
	pb "github.com/nurtilekassankhan/meetup/gateway/pkg/proto/accounting"
)

func (c *connection) RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error {

	_, err := c.conn.RequestWithdraw(ctx, &pb.RequestWithdrawRequest{
		Currency: data.Currency,
		Chain:    data.Chain,
		Amount:   data.Amount,
		Address:  data.Address,
	})
	if err != nil {
		return err
	}

	return nil

}
