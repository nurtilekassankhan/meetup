package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	RequestWithdraw(ctx *fiber.Ctx) error
	ConfirmWithdraw(ctx *fiber.Ctx) error
}

type handler struct {
}

func New()
