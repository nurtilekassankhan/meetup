package repository

import (
	"github.com/nurtilekassankhan/meetup/gateway/pkg/repository/accountingconn"
)

type Repository struct {
	Accounting accountingconn.Connection
}

func NewRepository(conn accountingconn.Connection) *Repository {
	return &Repository{
		Accounting: conn,
	}
}
