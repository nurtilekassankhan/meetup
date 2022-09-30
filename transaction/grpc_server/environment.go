package grpc_server

import (
	"context"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/config"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/handler"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/repository"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/repository/sqlite"
	"github.com/nurtilekassankhan/meetup/transaction/pkg/service/transactionsrv"
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

	repo := repository.NewRepository(db)

	transaction, err := transactionsrv.New(*repo)
	if err != nil {
		return nil, err
	}

	handlr, err := handler.New(transaction)
	if err != nil {
		return nil, err
	}

	srv, err := NewServer(&conf.Server, handlr)
	if err != nil {
		return nil, err
	}

	return srv, nil

}
