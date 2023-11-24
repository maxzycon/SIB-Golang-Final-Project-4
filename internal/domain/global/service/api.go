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

	// ----- Task
	GetAllTask(ctx context.Context, user *authutil.UserClaims) (resp []*dto.TaskRow, err error)
	CreateTask(ctx context.Context, payload *dto.PayloadTask, user *authutil.UserClaims) (resp *dto.TaskCreteResponse, err error)
	UpdateTaskById(ctx context.Context, id int, payload *dto.PayloadUpdateTask) (resp *dto.TaskUpdateResponse, err error)
	UpdateTaskStatusById(ctx context.Context, id int, payload *dto.PayloadUpdateStatusTask) (resp *dto.TaskUpdateResponse, err error)
	UpdateTaskCategoryById(ctx context.Context, id int, payload *dto.PayloadUpdateCategoryTask) (resp *dto.TaskUpdateResponse, err error)
	DeleteTaskById(ctx context.Context, id int) (resp *int64, err error)
}
