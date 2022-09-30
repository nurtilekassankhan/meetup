package accountingsrv

import (
	"context"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/common"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/model"
	"time"
)

func (s *service) RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error {

	// todo: validate withdraw data

	code, err := common.GenerateOtp()
	if err != nil {
		return err
	}

	_, err = s.repo.WithdrawRequest.Create(&model.WithdrawRequestDatabase{
		Currency:   data.Currency,
		Blockchain: data.Blockchain,
		Amount:     data.Amount,
		Address:    data.Address,
		Otp:        code,
		CreatedAt:  time.Time{},
	})
	if err != nil {
		return err
	}

	return nil

}
