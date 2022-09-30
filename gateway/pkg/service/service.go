package service

import (
	"context"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/config"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Conn interface {
	RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error
	ConfirmWithdraw(ctx context.Context, data *model.WithdrawConfirm) error
}

type conn struct {
	conn pb.UserManagementClient
}

func New(data *config.AccountingServer) (Conn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcConn, err := grpc.Dial(data.Host+data.Port, opts...)
	if err != nil {
		return nil, err
	}

	client := pb.NewUserManagementClient(grpcConn)

	return &conn{
		conn: client,
	}, nil
}
