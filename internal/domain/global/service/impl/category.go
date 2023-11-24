package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/model"
	"gorm.io/gorm"
)

func (s *GlobalService) GetAllCategory(ctx context.Context, user *authutil.UserClaims) (resp []*dto.CategoryRow, err error) {
	data := make([]*model.Category, 0)

	if err = s.db.WithContext(ctx).Preload("Product").Model(&model.Category{}).Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.CategoryRow, 0)

	for _, v := range data {
		tmp := &dto.CategoryRow{
			ID:        v.ID,
			Type:      v.Type,
			UpdatedAt: v.UpdatedAt,
			CreatedAt: v.CreatedAt,
		}

		for _, t := range v.Product {
			tmp.Products = append(tmp.Products, dto.ProductCategoryRow{
				ID:        t.ID,
				Title:     t.Title,
				Price:     t.Price,
				Stock:     t.Stock,
				CreatedAt: t.CreatedAt,
				UpdatedAt: t.UpdatedAt,
			})
		}

		resp = append(resp, tmp)
	}

	return resp, nil
}

func (s *GlobalService) CreateCategory(ctx context.Context, payload *dto.PayloadCategory, user *authutil.UserClaims) (resp *dto.CategoryCreteResponse, err error) {
	entity := &model.Category{
		Type:              payload.Type,
		SoldProductAmount: 0,
	}
	if err = s.db.Create(&entity).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err create Category")
		return
	}

	resp = &dto.CategoryCreteResponse{
		ID:                entity.ID,
		Type:              entity.Type,
		SoldProductAmount: entity.SoldProductAmount,
		CreatedAt:         entity.CreatedAt,
	}
	return
}

func (s *GlobalService) UpdateCategoryById(ctx context.Context, id int, payload *dto.PayloadUpdateCategory) (resp *dto.CategoryUpdateResponse, err error) {
	entity := &model.Category{
		Model: gorm.Model{
			ID: uint(id),
		},
		Type: payload.Type,
	}
	if err = s.db.Updates(&entity).Error; err != nil {
		return
	}

	data := &model.Category{}
	if err = s.db.Find(&data, id).Error; err != nil {
		return
	}

	resp = &dto.CategoryUpdateResponse{
		ID:                data.ID,
		Type:              data.Type,
		SoldProductAmount: data.SoldProductAmount,
		UpdatedAt:         data.UpdatedAt,
	}
	return
}

func (s *GlobalService) DeleteCategoryById(ctx context.Context, id int) (resp *int64, err error) {
	if err = s.db.Delete(&model.Category{}, id).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err delete Category %d", id)
		return
	}
	return
}
