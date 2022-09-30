package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/model"
	"github.com/nurtilekassankhan/meetup/libs/pkg/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) ConfirmWithdraw(ctx *fiber.Ctx) error {
	// здесь типа отрабатывает сервис auth
	isAuthenticated := true
	if !isAuthenticated {
		err := status.Error(codes.Unauthenticated, "invalid auth token") // ошибка, которую вернул сервис auth
		return common.SendJSONResponseWithGrpc(ctx, err)
	}

	body := &model.WithdrawConfirm{}
	if err := ctx.BodyParser(body); err != nil {
		return common.SendJSONResponseWithGrpc(ctx, err)
	}
	if err := h.services.ConfirmWithdraw(ctx.Context(), body); err != nil {
		return common.SendJSONResponseWithGrpc(ctx, err)
	}
	return common.SendJSONResponseWithGrpc(ctx, nil)
}
