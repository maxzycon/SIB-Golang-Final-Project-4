package service

import (
	"context"

	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/user/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	response "github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/dto"

	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/model"
)

type UserService interface {
	// ---- Users
	GetUserByIdToken(ctx context.Context, userId uint) (resp *model.User, err error)
	UpdateUserProfile(ctx context.Context, id int, password string) (resp *int64, err error)
	GetById(ctx context.Context, id int) (resp *dto.UserRowDetail, err error)
	DeleteUserById(ctx context.Context, claims *authutil.UserClaims) (resp response.BaseResponseMessage, err error)
	Login(ctx context.Context, payload dto.PayloadLogin) (resp *dto.LoginRes, err error)
	GetUserByUsername(ctx context.Context, username string) (resp *model.User, err error)

	UpdateUser(ctx context.Context, payload dto.PayloadUpdateUser, claims *authutil.UserClaims) (resp *dto.UserRowDetailUpdateRes, err error)
	CreateUser(ctx context.Context, payload dto.PayloadCreateUser) (resp *dto.UserRowDetail, err error)
}
