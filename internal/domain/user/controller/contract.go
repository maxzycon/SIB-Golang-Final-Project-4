package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/config"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/internal/domain/user/service"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/constant/role"
	"github.com/maxzycon/SIB-Golang-Final-Project-4/pkg/middleware"
)

const (
	Create = "/users/register"

	Update            = "/users/update-account"
	GetById           = "/users/:id"
	Delete            = "/users/delete-account"
	GetUsersPaginated = "/users"
	UpdateUserProfile = "/profile"

	Login        = "/users/login"
	GetUserLogin = "/users"
)

type UsersControllerParams struct {
	V1          fiber.Router
	Conf        *config.Config
	UserService service.UserService
	Middleware  middleware.GlobalMiddleware
}
type usersController struct {
	v1          fiber.Router
	conf        *config.Config
	userService service.UserService
	middleware  middleware.GlobalMiddleware
}

func New(params *UsersControllerParams) *usersController {
	return &usersController{
		v1:          params.V1,
		conf:        params.Conf,
		userService: params.UserService,
		middleware:  params.Middleware,
	}
}
func (pc *usersController) Init() {
	pc.v1.Get(GetUserLogin, pc.middleware.Protected([]uint{role.ROLE_ADMIN, role.ROLE_MEMBER}), pc.handlerUser)
	pc.v1.Get(GetById, pc.middleware.Protected([]uint{role.ROLE_ADMIN, role.ROLE_MEMBER}), pc.handlerGetUserById)
	pc.v1.Put(UpdateUserProfile, pc.middleware.Protected([]uint{role.ROLE_ADMIN, role.ROLE_MEMBER}), pc.handerUpdateUserProfile)

	pc.v1.Put(Update, pc.middleware.Protected([]uint{role.ROLE_ADMIN, role.ROLE_MEMBER}), pc.handlerUpdateUser)
	pc.v1.Delete(Delete, pc.middleware.Protected([]uint{role.ROLE_ADMIN, role.ROLE_MEMBER}), pc.handlerDeleteUserById)
	pc.v1.Post(Create, pc.handlerCreateUser)
	pc.v1.Post(Login, pc.handlerLogin)
}
