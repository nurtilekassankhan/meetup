package grpc_server

import (
	"context"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/config"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/handler"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/repository"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/repository/sqlite"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/repository/transactionconn"
	"github.com/nurtilekassankhan/meetup/accounting/pkg/service/accountingsrv"
)

func New(ctx context.Context, conf *config.Config) (Server, error) {

	db, err := sqlite.NewSQLiteDB()
	if err != nil {
		return nil, err
	}

	err = sqlite.RunMigrateScripts(db)
	if err != nil {
		return nil, err
	}

	transaction, err := transactionconn.New(&conf.Transaction)
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(db, transaction)

	accounting, err := accountingsrv.New(*repo)
	if err != nil {
		return nil, err
	}

	handlr, err := handler.New(accounting)
	if err != nil {
		return nil, err
	}

	srv, err := NewServer(&conf.Server, handlr)
	if err != nil {
		return nil, err
	}

	return srv, nil

}
