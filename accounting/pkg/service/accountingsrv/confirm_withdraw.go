package accountingsrv

import (
	"context"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/common"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/model"
)

func (s *service) ConfirmWithdraw(ctx context.Context, otp string) error {

	if common.IsValidOtp(otp) {
		resp, err := s.repo.WithdrawRequest.GetByOtpCode(otp)
		if err != nil {
			return err
		}
		err = s.repo.Transaction.RecordWithdraw(ctx, &model.WithdrawRequestDatabase{
			Currency:   resp.Currency,
			Blockchain: resp.Blockchain,
			Amount:     resp.Amount,
			Address:    resp.Address,
		})
		if err != nil {
			return err
		}

	} else {
		return accounting_errors
	}

	return nil

}
