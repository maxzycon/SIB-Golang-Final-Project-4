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

func (c *GlobalController) handlerCreateTask(f *fiber.Ctx) (err error) {
	payload := dto.PayloadTask{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Error("err parse body create Task")
		return httputil.WriteErrorResponse(f, err)
	}

	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.CreateTask(f.Context(), &payload, user)

	if err != nil {
		log.Errorf("err service at controller create Task :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusCreated, resp)
}

func (c *GlobalController) handlerUpdateTask(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Task")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateTask{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Task")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateTaskById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Task :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerUpdateStatusTask(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Task")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateStatusTask{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Task")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateTaskStatusById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Task :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerUpdateCategoryTask(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params update Task")
		return httputil.WriteErrorResponse(f, err)
	}

	payload := dto.PayloadUpdateCategoryTask{}
	err = f.BodyParser(&payload)
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse body update Task")
		return httputil.WriteErrorResponse(f, err)
	}
	resp, err := c.globalService.UpdateTaskCategoryById(f.Context(), id, &payload)

	if err != nil {
		log.Errorf("err service at controller update Task :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerGetAllTask(f *fiber.Ctx) (err error) {
	user, _ := authutil.GetCredential(f)
	resp, err := c.globalService.GetAllTask(f.Context(), user)

	if err != nil {
		log.Errorf("err service at controller Task :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseWriteResponse(f, http.StatusOK, resp)
}

func (c *GlobalController) handlerDeleteTask(f *fiber.Ctx) (err error) {
	id, err := f.ParamsInt("id")
	if err != nil {
		err = errors.ErrBadRequest
		log.Errorf("err parse params Task delete by id")
		return httputil.WriteErrorResponse(f, err)
	}
	_, err = c.globalService.DeleteTaskById(f.Context(), id)

	if err != nil {
		log.Errorf("err service at controller Task delete by id :%+v", err)
		return httputil.WriteErrorResponse(f, err)
	}

	return httputil.BaseMessageResponse(f, "Task has been successfully deleted")
}
