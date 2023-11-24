package impl

import (
	"context"
	"strings"
	"time"

	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/authutil"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/constant/role"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/user/dto"
	response "github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/dto"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/errors"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/helper"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/model"
)

func (s *UserService) Login(ctx context.Context, payload dto.PayloadLogin) (resp *dto.LoginRes, err error) {
	user, err := s.UserRepository.FindUserByEmailLogin(ctx, payload.Email)
	if err != nil {
		log.Errorf("[Login] findUserByusername :%+v", err)
		return
	}

	password := helper.CheckPasswordHash(payload.Password, user.Password)
	if !password {
		log.Errorf("[Login] err hash password doens't match")
		err = errors.ErrInvalidPassword
		return
	}

	// --- set 30 day exp
	exp := time.Now().Add((time.Hour * 24) * 30).Unix()
	claims := jwt.MapClaims{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
		"exp":       exp,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	AccessToken, err := tokenClaims.SignedString([]byte(s.conf.JWT_SECRET_KEY))

	if err != nil {
		log.Errorf("[Login] err create access token %+v", err)
		return
	}

	resp = &dto.LoginRes{
		AccessToken: AccessToken,
	}
	return
}

func (s *UserService) CreateUser(ctx context.Context, payload dto.PayloadCreateUser) (resp *dto.UserRowDetail, err error) {
	password, _ := helper.HashPassword(strings.Trim(payload.Password, " "))
	userPayload := model.User{
		FullName: payload.FullName,
		Email:    payload.Email,
		Password: password,
		Role:     role.ROLE_MEMBER,
	}

	data, err := s.UserRepository.Create(ctx, &userPayload)
	if err != nil {
		log.Errorf("[user.go][CreateUser] err create user :%+v", err)
		return
	}

	resp = &dto.UserRowDetail{
		ID:        data.ID,
		FullName:  data.FullName,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
	}

	return
}

func (s *UserService) GetById(ctx context.Context, id int) (resp *dto.UserRowDetail, err error) {
	row, err := s.UserRepository.FindById(ctx, id)
	if err != nil {
		log.Errorf("[user.go][GetById] err repository at service :%+v", err)
		return
	}
	resp = &dto.UserRowDetail{
		FullName:  row.FullName,
		Email:     row.Email,
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
	}
	return
}

func (s *UserService) UpdateUser(ctx context.Context, payload dto.PayloadUpdateUser, claims *authutil.UserClaims) (resp *dto.UserRowDetailUpdateRes, err error) {
	_, err = s.UserRepository.FindById(ctx, int(claims.ID))
	if err != nil {
		log.Errorf("[user.go][CreateUser] err create user :%+v", err)
		return
	}
	userPayload := model.User{
		FullName: payload.FullName,
		Email:    payload.Email,
	}

	data, err := s.UserRepository.Update(ctx, &userPayload, int(claims.ID))
	if err != nil {
		log.Errorf("[user.go][UpdateUser] err create user :%+v", err)
		return
	}

	resp = &dto.UserRowDetailUpdateRes{
		ID:        data.ID,
		Email:     data.Email,
		FullName:  data.FullName,
		UpdatedAt: data.UpdatedAt,
	}

	return
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (resp *model.User, err error) {
	resp, err = s.UserRepository.FindUserByUsername(ctx, username)
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}
	return
}

func (s *UserService) DeleteUserById(ctx context.Context, claims *authutil.UserClaims) (resp response.BaseResponseMessage, err error) {
	_, err = s.UserRepository.DeleteUserById(ctx, int(claims.ID))
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}

	return response.BaseResponseMessage{
		Message: "Your account has been succesfully deleted",
	}, nil
}

func (s *UserService) UpdateUserProfile(ctx context.Context, id int, password string) (resp *int64, err error) {
	newPassword, _ := helper.HashPassword(strings.Trim(password, " "))
	resp, err = s.UserRepository.UpdatePasswordByUserId(ctx, id, &newPassword)
	if err != nil {
		log.Errorf("[user.go][FindByUsername] err repository at service :%+v", err)
		return
	}
	return
}

func (s *UserService) GetUserByIdToken(ctx context.Context, userId uint) (resp *model.User, err error) {
	resp, err = s.UserRepository.GetUserByIdToken(ctx, userId)
	if err != nil {
		log.Errorf("[user.go][GetById] err repository at service :%+v", err)
		return
	}
	return
}
