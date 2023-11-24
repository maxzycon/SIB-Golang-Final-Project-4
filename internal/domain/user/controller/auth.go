package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/user/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/httputil"
)

func (c *usersController) handlerLogin(f *fiber.Ctx) (err error) {
	payload := dto.PayloadLogin{}
	err = f.BodyParser(&payload)
	if err != nil {
		log.Errorf("[handlerLogin] err body parse")
		err = errors.ErrBadRequest
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.userService.Login(f.Context(), payload)

	if err != nil {
		log.Errorf("[handlerLogin] err service at controller %v", err)
		return httputil.WriteErrorResponse(f, err)
	}
	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}
