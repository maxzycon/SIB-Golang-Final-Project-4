package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/user/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/httputil"
)

func (c *usersController) handlerUser(f *fiber.Ctx) (err error) {
	user, err := authutil.GetCredential(f)
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}
	return httputil.WriteSuccessResponse(f, user)
}

func (c *usersController) handerUpdateUserProfile(f *fiber.Ctx) (err error) {
	payload := dto.PayloadUpdateProfile{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handleUpdateUserProfile] err parse body")
		return httputil.WriteErrorResponse(f, err)
	}

	user, err := authutil.GetCredential(f)
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.userService.UpdateUserProfile(f.Context(), int(user.ID), payload.Password)
	if err != nil {
		log.Error("err update user profile controller")
		return httputil.WriteErrorResponse(f, err)
	}
	return httputil.WriteSuccessResponseAffectedRow(f, resp)
}

func (c *usersController) handlerCreateUser(f *fiber.Ctx) (err error) {
	payload := dto.PayloadCreateUser{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handleCreateUser] err parse body")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.userService.CreateUser(f.Context(), payload)

	if err != nil {
		log.Errorf("[user.go][hadnlerCreateUser] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusCreated, resp)
}

func (c *usersController) handlerUpdateUser(f *fiber.Ctx) (err error) {
	payload := dto.PayloadUpdateUser{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handleCreateUser] err parse body")
		return httputil.WriteErrorResponse(f, err)
	}

	user, err := authutil.GetCredential(f)
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.userService.UpdateUser(f.Context(), payload, user)

	if err != nil {
		log.Errorf("[user.go][hadnlerCreateUser] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *usersController) handlerGetUserById(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("[handlerGetUserById] err parse params")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.userService.GetById(f.Context(), id)

	if err != nil {
		log.Errorf("[user.go][handlerGetUserById] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.WriteSuccessResponse(f, resp)
}

func (c *usersController) handlerDeleteUserById(f *fiber.Ctx) (err error) {
	user, err := authutil.GetCredential(f)
	if err != nil {
		log.Errorf("err parse user")
		return httputil.WriteErrorResponse(f, err)
	}

	resp, err := c.userService.DeleteUserById(f.Context(), user)

	if err != nil {
		log.Errorf("[user.go][handlerDeleteUserById] err service at repo :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}
