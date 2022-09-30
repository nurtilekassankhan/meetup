package common

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/nurtilekassankhan/meetup/libs/pkg/accounting_errors"
	"github.com/nurtilekassankhan/meetup/libs/pkg/transaction_errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type StdResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	//Data    interface{} `json:"data"`
}

func SendJSONResponseWithoutGrpc(ctx *fiber.Ctx, response *StdResponse) error {

	ctx = ctx.Status(response.Code)

	return ctx.JSON(&StdResponse{
		Code:    response.Code,
		Message: response.Message,
	})

}

func SendJSONResponseWithGrpc(ctx *fiber.Ctx, err error) error {

	var (
		code int
		msg  string
	)

	st, ok := status.FromError(err)
	if !ok {
		if err == nil {
			code = http.StatusOK
			msg = "ok"
		} else {
			code = http.StatusInternalServerError
			msg = "unknown error" // todo: потом заменить просто на unknown error
		}
	}

	code = ToHTTPCode(st.Code().String())
	msg = st.Message()

	ctx = ctx.Status(code)

	return ctx.JSON(&StdResponse{
		Code:    code,
		Message: msg,
	})

}

func ToHTTPCode(statusCode string) int {
	switch statusCode {
	case codes.OK.String():
		return http.StatusOK
	case codes.InvalidArgument.String():
		return http.StatusBadRequest
	case codes.NotFound.String():
		return http.StatusNotFound
	case codes.PermissionDenied.String():
		return http.StatusForbidden
	case codes.Unimplemented.String():
		return http.StatusNotImplemented
	case codes.Internal.String():
		return http.StatusInternalServerError
	case codes.Unavailable.String():
		return http.StatusServiceUnavailable
	case codes.Unauthenticated.String():
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError // todo: возможно нужно будет заменить
	}
}

func ClassifyErrorToGRPCStatus(err error) (codes.Code, string) {

	// HTTP - 200
	if err == nil {
		return codes.OK, "ok"
	}

	// HTTP - 400
	if errors.Is(err, accounting_errors.InvalidAmount) {
		return codes.InvalidArgument, ""
	} else if errors.Is(err, transaction_errors.InvalidTransactionId) {
		return codes.InvalidArgument, ""
	}

	// HTTP - 500

	return codes.Internal, "unknown error"
}
