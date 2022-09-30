package transactionconn

import (
	"context"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/config"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/model"
	pb "github.com/nurtilekassankhan/meetup/accounting/pkg/proto/transaction"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Connection interface {
	RecordWithdraw(ctx context.Context, data *model.WithdrawRequestDatabase) error
}

type connection struct {
	conn pb.UserManagementClient
}

func New(conf *config.Transaction) (Connection, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(conf.Host+conf.Port, opts...)
	if err != nil {
		return nil, err
	}
	client := pb.NewUserManagementClient(conn)

	return &connection{conn: client}, nil
}
