package accountingconn

import (
	"context"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/config"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/model"
	pb "github.com/nurtilekassankhan/meetup/gateway/pkg/proto/accounting"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection interface {
	RequestWithdraw(ctx context.Context, data *model.WithdrawRequest) error
	ConfirmWithdraw(ctx context.Context, data *model.WithdrawConfirm) error
}

type connection struct {
	conn pb.UserManagementClient
}

func New(conf *config.AccountingServer) (Connection, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(conf.Host+conf.Port, opts...)
	if err != nil {
		return nil, err
	}
	client := pb.NewUserManagementClient(conn)

	return &connection{conn: client}, nil
}
