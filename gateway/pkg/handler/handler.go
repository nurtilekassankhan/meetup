package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/service"
)

type Handler interface {
	RequestWithdraw(ctx *fiber.Ctx) error
	ConfirmWithdraw(ctx *fiber.Ctx) error
}

type handler struct {
	services service.Service
}

func New(services service.Service) (Handler, error) {
	return &handler{
		services: services,
	}, nil
}
