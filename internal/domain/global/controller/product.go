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

func (c *GlobalController) handlerCreateProduct(f *fiber.Ctx) (err error) {
	payload := dto.PayloadProduct{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Error("err parse body create Product")
		return httputil.WriteErrorResponse(f, err)
	}

	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.CreateProduct(f.Context(), &payload, user)

	if err != nil {
		log.Errorf("err service at controller create Product :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusCreated, resp)
}

func (c *GlobalController) handlerUpdateProduct(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Product")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateProduct{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Product")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateProductById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Product :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, dto.ProductUpdateWrappper{
		Products: resp,
	})
}

func (c *GlobalController) handlerGetAllProduct(f *fiber.Ctx) (err error) {
	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.GetAllProduct(f.Context(), user)

	if err != nil {
		log.Errorf("err service at controller Product :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerDeleteProduct(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Product delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	_, err = c.globalService.DeleteProductById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Product delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseMessageResponse(f, "Product has been successfully deleted")
}
