package controller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/httputil"
)

func (c *GlobalController) handlerCreateCategory(f *fiber.Ctx) (err error) {
	payload := dto.PayloadCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Error("err parse body create Category")
		return httputil.WriteErrorResponse(f, err)
	}

	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.CreateCategory(f.Context(), &payload, user)

	if err != nil {
		log.Errorf("err service at controller create Category :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusCreated, resp)
}

func (c *GlobalController) handlerUpdateCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Category")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateCategory{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Category")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateCategoryById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Category :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerGetAllCategory(f *fiber.Ctx) (err error) {
	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.GetAllCategory(f.Context(), user)

	if err != nil {
		log.Errorf("err service at controller Category :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerDeleteCategory(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Category delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	_, err = c.globalService.DeleteCategoryById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Category delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseMessageResponse(f, "Category has been successfully deleted")
}
