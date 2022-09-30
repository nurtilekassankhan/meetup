package accountingconn

import (
	"context"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/model"
	pb "github.com/nurtilekassankhan/meetup/gateway/pkg/proto/accounting"
)

func (c *connection) ConfirmWithdraw(ctx context.Context, data *model.WithdrawConfirm) error {

	_, err := c.conn.ConfirmWithdraw(ctx, &pb.ConfirmWithdrawRequest{Code: data.OtpCode})
	if err != nil {
		return err
	}

	return err
}
