package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurtilekassankhan/meetup/gateway/pkg/model"
	"github.com/nurtilekassankhan/meetup/libs/pkg/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func (h *handler) RequestWithdraw(ctx *fiber.Ctx) error {
	// здесь типа отрабатывает сервис auth
	isAuthenticated := true
	if !isAuthenticated {
		err := status.Error(codes.Unauthenticated, "invalid auth token or refresh") // ошибка, которую вернул сервис auth
		return common.SendJSONResponseWithGrpc(ctx, err)
	}

	body := &model.WithdrawRequest{}
	if err := ctx.BodyParser(body); err != nil {
		// todo: logger
		return common.SendJSONResponseWithoutGrpc(ctx, &common.StdResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid body",
		})
	}

	if err := h.services.RequestWithdraw(ctx.Context(), &model.WithdrawRequest{
		Currency: body.Currency,
		Chain:    body.Chain,
		Amount:   body.Amount,
		Address:  body.Address,
	}); err != nil {
		// todo: logger
		return common.SendJSONResponseWithGrpc(ctx, err)
	}

	return common.SendJSONResponseWithGrpc(ctx, nil)
	
}
