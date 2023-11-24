package impl

import (
	"context"
	"fmt"

	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/model"
	"gorm.io/gorm"
)

func (s *GlobalService) GetAllMyTransaction(ctx context.Context, user *authutil.UserClaims) (resp []*dto.MyTransaction, err error) {
	data := make([]*model.TransactionHistory, 0)

	if err = s.db.WithContext(ctx).Preload("Product").Model(&model.TransactionHistory{}).Where("user_id = ?", user.ID).Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.MyTransaction, 0)

	for _, v := range data {
		tmp := &dto.MyTransaction{
			ID:         v.ID,
			ProductID:  v.ProductID,
			UserID:     v.UserID,
			Quantity:   v.Quantity,
			TotalPrice: v.Quantity,
			Product: dto.ProductMyTransaction{
				ID:         v.ProductID,
				Title:      v.Product.Title,
				Price:      v.Product.Price,
				Stock:      v.Product.Stock,
				CategoryID: v.Product.CategoryID,
				CreatedAt:  v.Product.CreatedAt,
				UpdatedAt:  v.UpdatedAt,
			},
		}

		resp = append(resp, tmp)
	}

	return resp, nil
}

func (s *GlobalService) GetAllUserTransaction(ctx context.Context, user *authutil.UserClaims) (resp []*dto.UserTransaction, err error) {
	data := make([]*model.TransactionHistory, 0)

	if err = s.db.WithContext(ctx).Preload("Product").Preload("User").Model(&model.TransactionHistory{}).Find(&data).Error; err != nil {
		return
	}

	resp = make([]*dto.UserTransaction, 0)

	for _, v := range data {
		tmp := &dto.UserTransaction{
			ID:         v.ID,
			ProductID:  v.ProductID,
			UserID:     v.UserID,
			Quantity:   v.Quantity,
			TotalPrice: v.Quantity,
			Product: dto.ProductUserTransaction{
				ID:         v.ProductID,
				Title:      v.Product.Title,
				Price:      v.Product.Price,
				Stock:      v.Product.Stock,
				CategoryID: v.Product.CategoryID,
				CreatedAt:  v.Product.CreatedAt,
				UpdatedAt:  v.UpdatedAt,
			},
			User: dto.ProductUserInfo{
				ID:        v.UserID,
				Email:     v.User.Email,
				FullName:  v.User.FullName,
				Balance:   v.User.Balance,
				CreatedAt: v.User.CreatedAt,
				UpdatedAt: v.User.UpdatedAt,
			},
		}

		resp = append(resp, tmp)
	}

	return resp, nil
}

func (s *GlobalService) CreateTransaction(ctx context.Context, payload *dto.PayloadTransaction, user *authutil.UserClaims) (resp *dto.ResponseCreateTransaction, err error) {
	var totalPrice uint64 = 0
	var productTitle = ""
	// --- db transcation
	err = s.db.Transaction(func(tx *gorm.DB) (err error) {
		// ----- Check product exist or not
		product := &model.Product{}
		if err = tx.Model(&model.Product{}).Preload("Category").First(product, payload.ProductID).Error; err != nil {
			return
		}

		// ----- chceck stock > qty
		if product.Stock < payload.Quantity {
			err = errors.ErrInvalidStock
			return
		}

		// ----- check balance > price * qty
		if user.Balance < product.Price*payload.Quantity {
			err = errors.ErrInvalidBalance
			return
		}

		// ----- Update stock
		if err = tx.Model(&model.Product{}).Where("id = ?", product.ID).Updates(&model.Product{
			Stock: product.Stock - payload.Quantity,
		}).Error; err != nil {
			return
		}

		// ----- Update user balance
		if err = tx.Model(&model.User{}).Where("id = ?", user.ID).Updates(&model.User{
			Balance: user.Balance - (product.Price * payload.Quantity),
		}).Error; err != nil {
			return
		}

		// ----- Update sold category item
		if err = tx.Model(&model.Category{}).Where("id = ?", product.CategoryID).Updates(&model.Category{
			SoldProductAmount: product.Category.SoldProductAmount + payload.Quantity,
		}).Error; err != nil {
			return
		}

		entity := &model.TransactionHistory{
			ProductID:  uint(payload.ProductID),
			UserID:     user.ID,
			Quantity:   payload.Quantity,
			TotalPrice: product.Price * payload.Quantity,
		}

		if err = tx.Create(&entity).Error; err != nil {
			fmt.Println(err)
			return
		}
		totalPrice = product.Price * payload.Quantity
		productTitle = product.Title

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return
	}

	resp = &dto.ResponseCreateTransaction{
		Message: "You have succesfully purchased the order",
		TransactionBill: dto.TransactionBill{
			Quantity:     payload.Quantity,
			TotalPrice:   totalPrice,
			ProductTitle: productTitle,
		},
	}
	return
}
