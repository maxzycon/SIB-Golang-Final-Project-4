package impl

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/model"
	"gorm.io/gorm"
)

func (s *GlobalService) GetAllProduct(ctx context.Context, user *authutil.UserClaims) (resp []*dto.ProductRow, err error) {
	data := make([]*model.Product, 0)

	if err = s.db.WithContext(ctx).Preload("User").Model(&model.Product{}).Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.ProductRow, 0)

	for _, v := range data {
		tmp := &dto.ProductRow{
			ID:         v.ID,
			Title:      v.Title,
			CategoryID: v.CategoryID,
			Price:      v.Price,
			Stock:      v.Stock,
			CreatedAt:  v.CreatedAt,
		}

		resp = append(resp, tmp)
	}

	return resp, nil
}

func (s *GlobalService) CreateProduct(ctx context.Context, payload *dto.PayloadProduct, user *authutil.UserClaims) (resp *dto.ProductCreteResponse, err error) {
	entity := &model.Product{
		Title:      payload.Title,
		Price:      payload.Price,
		Stock:      payload.Stock,
		CategoryID: payload.CategoryID,
	}
	if err = s.db.Create(&entity).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err create Product")
		return
	}

	resp = &dto.ProductCreteResponse{
		ID:         entity.ID,
		Title:      entity.Title,
		Price:      entity.Price,
		Stock:      entity.Stock,
		CategoryID: payload.CategoryID,
		CreatedAt:  entity.CreatedAt,
	}
	return
}

func (s *GlobalService) UpdateProductById(ctx context.Context, id int, payload *dto.PayloadUpdateProduct) (resp *dto.ProductUpdateResponse, err error) {
	entity := &model.Product{
		Model: gorm.Model{
			ID: uint(id),
		},
		Title:      payload.Title,
		Price:      payload.Price,
		Stock:      payload.Stock,
		CategoryID: payload.CategoryID,
	}
	if err = s.db.Updates(&entity).Error; err != nil {
		return
	}

	data := &model.Product{}
	if err = s.db.Find(&data, id).Error; err != nil {
		return
	}

	resp = &dto.ProductUpdateResponse{
		ID:         data.ID,
		Title:      entity.Title,
		Price:      entity.Price,
		Stock:      data.Stock,
		CreatedAt:  data.CreatedAt,
		CategoryID: data.CategoryID,
		UpdatedAt:  data.UpdatedAt,
	}
	return
}

func (s *GlobalService) DeleteProductById(ctx context.Context, id int) (resp *int64, err error) {
	if err = s.db.Delete(&model.Product{}, id).Error; err != nil {
		return
	}

	if err != nil {
		log.Errorf("err delete Product %d", id)
		return
	}
	return
}
