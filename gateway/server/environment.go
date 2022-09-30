package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/config"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/repository"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/repository/accountingconn"
)

func New(conf *config.Config) (Server, error) {

	conn, err := accountingconn.New(&conf.Accounting)
	if err != nil {
		return nil, err
	}

	repo := repository.NewRepository(conn)

	services, err := NewServices(*repo)
	if err != nil {
		return nil, err
	}
	controllers, err := NewHandlers(services)
	if err != nil {
		return nil, err
	}
	app := fiber.New()

	server, err := NewServer(conf.Server.Port, app, controllers)
	if err != nil {
		return nil, err
	}

	return server, nil

}