package server

import (
	"github.com/nurtilekassankhan/meetup/gateway/pkg/repository"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/service"
)

type servs struct {
	service.Service
}

func NewServices(repo repository.Repository) (*servs, error) {
	srvs := service.New(repo)
	return &servs{srvs}, nil
}
