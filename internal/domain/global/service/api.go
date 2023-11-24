package service

import (
	"context"

	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/global/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
)

type GlobalService interface {
	// ----- Categorys
	GetAllCategory(ctx context.Context, user *authutil.UserClaims) (resp []*dto.CategoryRow, err error)
	CreateCategory(ctx context.Context, payload *dto.PayloadCategory, user *authutil.UserClaims) (resp *dto.CategoryCreteResponse, err error)
	UpdateCategoryById(ctx context.Context, id int, payload *dto.PayloadUpdateCategory) (resp *dto.CategoryUpdateResponse, err error)
	DeleteCategoryById(ctx context.Context, id int) (resp *int64, err error)

	// ----- Product
	GetAllProduct(ctx context.Context, user *authutil.UserClaims) (resp []*dto.ProductRow, err error)
	CreateProduct(ctx context.Context, payload *dto.PayloadProduct, user *authutil.UserClaims) (resp *dto.ProductCreteResponse, err error)
	UpdateProductById(ctx context.Context, id int, payload *dto.PayloadUpdateProduct) (resp *dto.ProductUpdateResponse, err error)
	DeleteProductById(ctx context.Context, id int) (resp *int64, err error)

	// ----- Transaction
	CreateTransaction(ctx context.Context, payload *dto.PayloadTransaction, user *authutil.UserClaims) (resp *dto.ResponseCreateTransaction, err error)
	GetAllUserTransaction(ctx context.Context, user *authutil.UserClaims) (resp []*dto.UserTransaction, err error)
	GetAllMyTransaction(ctx context.Context, user *authutil.UserClaims) (resp []*dto.MyTransaction, err error)
}
