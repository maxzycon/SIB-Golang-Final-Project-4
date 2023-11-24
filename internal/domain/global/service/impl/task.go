package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/model"
	"gorm.io/gorm"
)

func (s *GlobalService) GetAllTask(ctx context.Context, user *authutil.UserClaims) (resp []*dto.TaskRow, err error) {
	data := make([]*model.Task, 0)

	if err = s.db.WithContext(ctx).Preload("User").Model(&model.Task{}).Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.TaskRow, 0)

	for _, v := range data {
		tmp := &dto.TaskRow{
			ID:          v.ID,
			Title:       v.Title,
			Status:      v.Status,
			Description: v.Description,
			UserID:      v.UserID,
			CategoryID:  v.CategoryID,
			CreatedAt:   v.CreatedAt,
			User: &dto.UserTaskRow{
				ID:       v.UserID,
				Email:    v.User.Email,
				FullName: v.User.FullName,
			},
		}

		resp = append(resp, tmp)
	}

	return resp, nil
}

func (s *GlobalService) CreateTask(ctx context.Context, payload *dto.PayloadTask, user *authutil.UserClaims) (resp *dto.TaskCreteResponse, err error) {
	entity := &model.Task{
		Title:       payload.Title,
		Description: payload.Description,
		Status:      false,
		UserID:      user.ID,
		CategoryID:  payload.CategoryID,
	}
	if err = s.db.Create(&entity).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err create Task")
		return
	}

	resp = &dto.TaskCreteResponse{
		ID:          entity.ID,
		Title:       entity.Title,
		Status:      entity.Status,
		Description: entity.Description,
		UserID:      user.ID,
		CategoryID:  payload.CategoryID,
		CreatedAt:   entity.CreatedAt,
	}
	return
}

func (s *GlobalService) UpdateTaskById(ctx context.Context, id int, payload *dto.PayloadUpdateTask) (resp *dto.TaskUpdateResponse, err error) {
	entity := &model.Task{
		Model: gorm.Model{
			ID: uint(id),
		},
		Title:       payload.Title,
		Description: payload.Description,
	}
	if err = s.db.Updates(&entity).Error; err != nil {
		return
	}

	data := &model.Task{}
	if err = s.db.Find(&data, id).Error; err != nil {
		return
	}

	resp = &dto.TaskUpdateResponse{
		ID:          data.ID,
		Title:       entity.Title,
		Status:      entity.Status,
		Description: entity.Description,
		UserID:      data.UserID,
		CategoryID:  data.CategoryID,
		UpdatedAt:   data.UpdatedAt,
	}
	return
}

func (s *GlobalService) UpdateTaskStatusById(ctx context.Context, id int, payload *dto.PayloadUpdateStatusTask) (resp *dto.TaskUpdateResponse, err error) {
	if err = s.db.Model(&model.Task{}).Where("id = ?", id).Update("status", payload.Status).Error; err != nil {
		return
	}

	data := &model.Task{}
	if err = s.db.Find(&data, id).Error; err != nil {
		return
	}

	resp = &dto.TaskUpdateResponse{
		ID:          data.ID,
		Title:       data.Title,
		Status:      data.Status,
		Description: data.Description,
		UserID:      data.UserID,
		CategoryID:  data.CategoryID,
		UpdatedAt:   data.UpdatedAt,
	}
	return
}

func (s *GlobalService) UpdateTaskCategoryById(ctx context.Context, id int, payload *dto.PayloadUpdateCategoryTask) (resp *dto.TaskUpdateResponse, err error) {
	if err = s.db.Model(&model.Task{}).Where("id = ?", id).Update("category_id", payload.CategoryID).Error; err != nil {
		return
	}

	data := &model.Task{}
	if err = s.db.Find(&data, id).Error; err != nil {
		return
	}

	resp = &dto.TaskUpdateResponse{
		ID:          data.ID,
		Title:       data.Title,
		Status:      data.Status,
		Description: data.Description,
		UserID:      data.UserID,
		CategoryID:  data.CategoryID,
		UpdatedAt:   data.UpdatedAt,
	}
	return
}

func (s *GlobalService) DeleteTaskById(ctx context.Context, id int) (resp *int64, err error) {
	if err = s.db.Delete(&model.Task{}, id).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err delete Task %d", id)
		return
	}
	return
}
