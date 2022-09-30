package server

import "github.com/nurtilekassankhan/meetup/gateway/pkg/handler"

type handlr struct {
	handler.Handler
}

func NewHandlers(services *servs) (*handlr, error) {
	handlers, err := handler.New(services)
	if err != nil {
		return nil, err
	}

	return &handlr{handlers}, nil
}
