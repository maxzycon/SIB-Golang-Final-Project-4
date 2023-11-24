package httputil

import (
	"net/http"

	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/errors"
)

func WriteSuccessResponse(e *fiber.Ctx, payload interface{}) error {
	return WriteResponse(e, dto.ResponseParam{
		Status: http.StatusOK,
		Payload: dto.BaseResponse{
			Data: payload,
		},
	})
}

func WriteSuccessResponseAffectedRow(e *fiber.Ctx, affectedRow *int64) error {
	if *affectedRow > 0 {
		return WriteResponse(e, dto.ResponseParam{
			Status: http.StatusOK,
			Payload: dto.BaseResponse{
				Data: constant.Success,
			},
		})
	}
	return WriteErrorResponse(e, errors.ErrInternalServer)
}

func WriteErrorResponse(e *fiber.Ctx, er error) error {
	errResp := errors.GetErrorResponse(er)
	return WriteResponse(e, dto.ResponseParam{
		Status: int(errResp.HTTPCode),
		Payload: dto.BaseResponse{
			Error: &dto.ErrorResponse{
				Code:    errResp.Code,
				Message: errResp.Message,
			},
		},
	})
}

func WriteResponse(e *fiber.Ctx, param dto.ResponseParam) error {
	return e.Status(param.Status).JSON(param.Payload)
}

func BaseWriteResponse(e *fiber.Ctx, status int, payload interface{}) error {
	return e.Status(status).JSON(payload)
}

func BaseMessageResponse(e *fiber.Ctx, message interface{}) error {
	return e.Status(http.StatusOK).JSON(dto.BaseResponseMessage{
		Message: message,
	})
}
